package modules

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	client_api "github.com/prometheus/client_golang/api"
	client_api_v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/rules"
)

// Manager ...
type Manager struct {
	Config  Config
	Prom    Prom
	Options *rules.ManagerOptions
	Manager *rules.Manager
	Rules   Rules

	logger log.Logger
}

// NewManager ...
func NewManager(ctx context.Context, logger log.Logger,
	prom Prom, config Config) (*Manager, error) {
	localStorage, err := NewMockStorage()
	if err != nil {
		return nil, err
	}

	options := &rules.ManagerOptions{
		Appendable: localStorage,
		TSDB:       localStorage,
		QueryFunc: HTTPQueryFunc(
			log.With(logger, "component", "http query func"),
			prom.URL,
		),
		NotifyFunc: HTTPNotifyFunc(
			log.With(logger, "component", "http notify func"),
			config.AuthToken,
			fmt.Sprintf("%s%s", config.GatewayURL, config.GatewayPathNotify),
			config.NotifyReties,
		),
		Context:         ctx,
		ExternalURL:     &url.URL{},
		Registerer:      nil,
		Logger:          log.With(logger, "component", "rule manager"),
		OutageTolerance: time.Hour,        // default 1h
		ForGracePeriod:  10 * time.Minute, // default 10m
		ResendDelay:     time.Minute,      // default 1m
	}
	manager := rules.NewManager(options)

	return &Manager{
		Config:  config,
		Prom:    prom,
		Options: options,
		Manager: manager,
		Rules:   Rules{},

		logger: logger,
	}, nil
}

// Update ...
func (m *Manager) Update(rules Rules) error {
	m.Rules = rules

	filepath := filepath.Join(os.TempDir(), fmt.Sprintf("rule.%d.yml", m.Prom.ID))

	content, err := rules.Content()
	if err != nil {
		level.Error(m.logger).Log("msg", "get prom rule error", "error", err, "prom_id", m.Prom.ID)
		return err
	}

	err = ioutil.WriteFile(filepath, content, 0644)
	if err != nil {
		level.Error(m.logger).Log("msg", "write file error", "error", err, "prom_id", m.Prom.ID)
		return err
	}

	return m.Manager.Update(time.Duration(m.Config.EvaluationInterval), []string{filepath}, nil)
}

// Run ...
func (m *Manager) Run() {
	level.Info(m.logger).Log("msg", "start rule manager", "prom_id", m.Prom.ID)
	m.Manager.Run()
}

// Stop ...
func (m *Manager) Stop() {
	level.Info(m.logger).Log("msg", "stop rule manager", "prom_id", m.Prom.ID)
	m.Manager.Stop()
}

// DebugNotifyFunc
func DebugNotifyFunc(logger log.Logger) rules.NotifyFunc {
	return func(ctx context.Context, expr string, alerts ...*rules.Alert) {
		for _, i := range alerts {
			level.Debug(logger).Log(
				"msg", "send alert",
				"state", i.State.String(),
				"annotations", i.Annotations.String(),
				"labels", i.Labels.String(),
			)
		}
	}
}

// Alert
type Alert rules.Alert

// MarshalJSON ...
func (a *Alert) MarshalJSON() ([]byte, error) {
	for idx, i := range a.Labels {
		if i.Name == "alertname" {
			a.Labels = append(a.Labels[:idx], a.Labels[idx+1:]...)
		}
	}
	return json.Marshal(map[string]interface{}{
		"state":        a.State,
		"labels":       a.Labels,
		"annotations":  a.Annotations,
		"value":        math.Round(a.Value*100) / 100,
		"active_at":    a.ActiveAt,
		"fired_at":     a.FiredAt,
		"resolved_at":  a.ResolvedAt,
		"last_sent_at": a.LastSentAt,
		"valid_until":  a.ValidUntil,
	})
}

// HTTPNotifyFunc
func HTTPNotifyFunc(logger log.Logger, token string, url string, retries int) rules.NotifyFunc {
	return func(ctx context.Context, expr string, alerts ...*rules.Alert) {
		if len(alerts) == 0 {
			return
		}

		new := []*Alert{}
		for _, i := range alerts {
			new = append(new, (*Alert)(i))
		}

		data, err := json.Marshal(new)
		if err != nil {
			level.Error(logger).Log("msg", "encode json error", "error", err, "alerts", alerts)
			return
		}
		level.Debug(logger).Log("msg", "encode alerts success", "json", data)

		for i := 1; i <= retries; i++ {
			client := http.Client{
				Timeout: 5 * time.Second, // FIXME: timeout
			}
			req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
			req.Header.Add("Token", token)
			req.Header.Add("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				level.Error(logger).Log("msg", "notify error", "url", url, "error", err, "retries", i)
				continue
			}
			if resp.StatusCode == 200 {
				level.Debug(logger).Log("msg", "notify success", "url", url)
				break
			}
			level.Error(logger).Log("msg", "notify error", "url", url, "status", resp.StatusCode, "retries", i)
		}
	}
}

// HTTPQueryFunc
// TODO: use http keep-alive
func HTTPQueryFunc(logger log.Logger, url string) rules.QueryFunc {
	client, _ := client_api.NewClient(client_api.Config{
		Address: url,
	})
	api := client_api_v1.NewAPI(client)
	return func(ctx context.Context, q string, t time.Time) (promql.Vector, error) {
		vector := promql.Vector{}

		value, _, err := api.Query(ctx, q, t)
		if err != nil {
			return vector, err
		}
		switch value.Type() {
		case model.ValVector:
			for _, i := range value.(model.Vector) {
				l := labels.Labels{}
				for k, v := range i.Metric {
					l = append(l, labels.Label{
						Name:  string(k),
						Value: string(v),
					})
				}
				vector = append(vector, promql.Sample{
					Point: promql.Point{
						T: int64(i.Timestamp),
						V: float64(i.Value),
					},
					Metric: l,
				})
			}
			level.Debug(logger).Log(
				"msg", "query vector seccess",
				"query", q,
				"vector", vector,
			)
			return vector, nil
		default:
			// TODO: other type: "matrix" | "vector" | "scalar" | "string",
			return vector, fmt.Errorf("unknown result type [%s] query=[%s]", value.Type().String(), q)
		}
	}
}
