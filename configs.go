package main

import (
	"github.com/saas-templates/go-svelte/api"
	"github.com/saas-templates/go-svelte/pkg/config"
	"github.com/saas-templates/go-svelte/pkg/log"
	"github.com/spf13/cobra"
)

type Config struct {
	API       api.Config `mapstructure:"api"`
	Addr      string     `mapstructure:"addr" default:":8080"`
	LogFormat string     `mapstructure:"log_format" default:"json"`
	LogLevel  string     `mapstructure:"log_level" default:"warn"`
}

func loadConf(cmd *cobra.Command) Config {
	opts := []config.Option{
		config.WithEnv(),
		config.WithName("config"),
		config.WithCommand(cmd),
	}

	var cfg Config
	if err := config.Load(&cfg, opts...); err != nil {
		log.Fatalf(cmd.Context(), "failed to load configs: %v", err)
	}

	log.Setup(cfg.LogLevel, cfg.LogFormat)
	return cfg
}
