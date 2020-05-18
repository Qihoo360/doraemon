package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/common/promlog"
	promlogflag "github.com/prometheus/common/promlog/flag"
	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/Qihoo360/doraemon/cmd/rule-engine/modules"
)

var (
	// Version ...
	Version string = "develop"
	// Commit ...
	Commit string = ""
	// BuildTime ...
	BuildTime string = ""
	// BuildUser ...
	BuildUser string = ""
)

func main() {

	cfg := modules.Config{
		PromlogConfig: promlog.Config{},
	}

	a := kingpin.New(filepath.Base(os.Args[0]), "Rule Engine")

	a.HelpFlag.Short('h')

	a.Flag("gateway.url", "alert gateway url").
		Default("http://localhost:8080").StringVar(&cfg.GatewayURL)
	a.Flag("gateway.path.rule", "alert gateway rule url").
		Default("/api/v1/rules").StringVar(&cfg.GatewayPathRule)
	a.Flag("gateway.path.prom", "alert gateway prom url").
		Default("/api/v1/proms").StringVar(&cfg.GatewayPathProm)
	a.Flag("gateway.path.notify", "alert gateway notify url").
		Default("/api/v1/alerts").StringVar(&cfg.GatewayPathNotify)

	a.Flag("notify.retries", "notify retries").
		Default("3").IntVar(&cfg.NotifyReties)
	a.Flag("evaluation.interval", "rule evaluation interval").
		Default("15s").SetValue(&cfg.EvaluationInterval)
	a.Flag("reload.interval", "rule reload interval").
		Default("5m").SetValue(&cfg.ReloadInterval)
	a.Flag("auth.token", "http auth token").
		Default("96smhbNpRguoJOCEKNrMqQ").StringVar(&cfg.AuthToken)

	promlogflag.AddFlags(a, &cfg.PromlogConfig)

	_, err := a.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrapf(err, "Error parsing commandline arguments"))
		a.Usage(os.Args[1:])
		os.Exit(2)
	}

	logger := promlog.New(&cfg.PromlogConfig)

	level.Info(logger).Log("version", Version, "commit", Commit, "build_time", BuildTime, "build_user", BuildUser)

	reloader := modules.NewReloader(logger, cfg)

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigs
		level.Info(logger).Log("msg", "receive signal, stoping", "signal", sig)
		reloader.Stop()
	}()

	reloader.Run()
	reloader.Loop()

	level.Info(logger).Log("msg", "See you next time!")
}
