package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/saas-templates/go-svelte/api"
	"github.com/saas-templates/go-svelte/pkg/httputils"
	"github.com/saas-templates/go-svelte/pkg/log"
	"github.com/saas-templates/go-svelte/ui"
)

func newCLI(ctx context.Context) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "go-svelte",
		Short:   "A short summary",
		Version: fmt.Sprintf("%s\ncommit: %s\nbuilt on: %s\n", Version, Commit, BuiltAt),
	}

	rootCmd.SetContext(ctx)
	rootCmd.PersistentFlags().StringP("config", "c", "", "Config file override")

	rootCmd.AddCommand(
		cmdServe(),
		cmdConfigs(),
	)

	return rootCmd
}

func cmdServe() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := loadConf(cmd)

			router := chi.NewRouter()
			router.Mount("/", http.FileServer(http.FS(ui.DistFS)))
			router.Mount("/api", api.Router(cfg.API))

			log.Infof(cmd.Context(), "listening at '%s'...", cfg.Addr)
			if err := httputils.GracefulServe(cmd.Context(), 5*time.Second, cfg.Addr, router); err != nil {
				log.Fatalf(cmd.Context(), "server exited: %v", err)
			}
			log.Infof(cmd.Context(), "server exited gracefully")
		},
	}

	return cmd
}

func cmdConfigs() *cobra.Command {
	return &cobra.Command{
		Use:   "configs",
		Short: "Display current loaded configs",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := loadConf(cmd)

			m := map[string]any{}
			if err := mapstructure.Decode(cfg, &m); err != nil {
				log.Fatalf(cmd.Context(), "failed to decode")
			}
			_ = yaml.NewEncoder(os.Stdout).Encode(m)
		},
	}
}
