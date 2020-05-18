package modules

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/common/model"
	"github.com/prometheus/common/promlog"
)

// Config ...
type Config struct {
	NotifyReties       int
	GatewayURL         string
	GatewayPathRule    string
	GatewayPathProm    string
	GatewayPathNotify  string
	EvaluationInterval model.Duration
	ReloadInterval     model.Duration
	AuthToken          string

	PromlogConfig promlog.Config
}

// Reloader ...
type Reloader struct {
	config   Config
	managers []*Manager
	logger   log.Logger
	context  context.Context
	cancel   context.CancelFunc
	running  bool
}

// NewReloader ...
func NewReloader(logger log.Logger, cfg Config) *Reloader {
	ctx, cancel := context.WithCancel(context.Background())

	reloader := Reloader{
		config: cfg,

		logger:  logger,
		context: ctx,
		cancel:  cancel,
		running: false,
	}

	return &reloader
}

// Run rule manager
func (r *Reloader) Run() {
	r.running = true
	for _, i := range r.managers {
		i.Run()
	}
}

// Stop rule manager
func (r *Reloader) Stop() {
	r.running = false
	r.cancel()
	for _, i := range r.managers {
		i.Stop()
	}

}

// download the rules and update rule manager
func (r *Reloader) Update() error {
	level.Debug(r.logger).Log("msg", "start update rule")

	promrules, err := r.getPromRules()
	if err != nil {
		return err
	}

	// stop invalid manager
	for idx, m := range r.managers {
		del := true
		for _, p := range promrules {
			if m.Prom.ID == p.Prom.ID && m.Prom.URL == p.Prom.URL && p.Prom.URL != "" {
				del = false
			}
		}
		if del {
			level.Info(r.logger).Log("msg", "prom not exist, delete manager", "prom_id", m.Prom.ID, "prom_url", m.Prom.URL)
			m.Stop()
			r.managers = append(r.managers[:idx], r.managers[idx+1:]...)
		}
	}

	// update rules
	for _, p := range promrules {
		if p.Prom.URL == "" {
			level.Error(r.logger).Log("msg", "prom url is null", "prom_id", p.Prom.ID, "prom_url", p.Prom.URL)
			continue
		}
		var manager *Manager
		for _, m := range r.managers {
			if m.Prom.ID == p.Prom.ID && m.Prom.URL == p.Prom.URL && p.Prom.URL != "" {
				manager = m
			}
		}
		if manager == nil {
			m, err := NewManager(r.context, r.logger, p.Prom, r.config)
			m.Run()
			manager = m
			r.managers = append(r.managers, manager)
			if err != nil {
				level.Error(r.logger).Log("msg", "create manager error", "error", err, "prom_id", manager.Prom.ID, "prom_url", manager.Prom.URL)
				return err
			}
		}
		if manager != nil {
			err := manager.Update(p.Rules)
			if err != nil {
				level.Error(r.logger).Log("msg", "update rule error", "error", err, "prom_id", manager.Prom.ID, "prom_url", manager.Prom.URL)
			} else {
				level.Info(r.logger).Log("msg", "update rule success", "len", len(p.Rules), "prom_id", manager.Prom.ID, "prom_url", manager.Prom.URL)
			}
		}
	}

	level.Debug(r.logger).Log("msg", "end update rule")
	return nil
}

// Loop for checking the rules
func (r *Reloader) Loop() {
	for r.running {
		r.Update()

		select {
		case <-r.context.Done():
		case <-time.After(time.Duration(r.config.ReloadInterval)):
		}
	}
}

func (r *Reloader) getPromRules() ([]PromRules, error) {
	data := []PromRules{}
	client := http.Client{
		Timeout: 10 * time.Second, // FIXME: timeout
	}
	url := fmt.Sprintf("%s%s", r.config.GatewayURL, r.config.GatewayPathRule)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Token", r.config.AuthToken)
	resp, err := client.Do(req)
	if err != nil {
		return data, err
	}

	ruleresp := RulesResp{}
	decodor := json.NewDecoder(resp.Body)
	err = decodor.Decode(&ruleresp)
	if err != nil {
		level.Error(r.logger).Log("msg", "decode rule error", "error", err)
		return data, err
	}

	if ruleresp.Code != 0 {
		err = fmt.Errorf("get rules error: %s", ruleresp.Msg)
		return data, err
	}
	level.Debug(r.logger).Log("msg", "get rule success", "len", len(ruleresp.Data))

	data = ruleresp.Data.PromRules()

	// get prom url
	url = fmt.Sprintf("%s%s", r.config.GatewayURL, r.config.GatewayPathProm)
	req, _ = http.NewRequest("GET", url, nil)
	req.Header.Add("Token", r.config.AuthToken)
	resp, err = client.Do(req)
	if err != nil {
		return data, err
	}
	promresp := PromsResp{}
	decodor = json.NewDecoder(resp.Body)
	err = decodor.Decode(&promresp)
	if err != nil {
		level.Error(r.logger).Log("msg", "decode prom error", "error", err)
		return data, err
	}

	if promresp.Code != 0 {
		err = fmt.Errorf("get prom error: %s", promresp.Msg)
		return data, err
	}
	level.Debug(r.logger).Log("msg", "get prom success", "len", len(promresp.Data))

	//fill in url
	for idx, i := range data {
		for _, j := range promresp.Data {
			if i.Prom.ID == j.ID {
				data[idx].Prom.URL = j.URL
				break
			}
		}
	}

	return data, nil
}
