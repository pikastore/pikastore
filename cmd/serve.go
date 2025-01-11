package cmd

import (
	"errors"
	"log"
	"net/http"

	"github.com/pikastore/pikastore/core"
	"github.com/spf13/cobra"
)

func ServeCmd() *cobra.Command {
	var addr string
	var httpsaddr string
	cmd := &cobra.Command{
		Use:          "serve",
		Args:         cobra.ArbitraryArgs,
		Short:        "Starts the HTTP server",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && addr == "" {
				if addr == "" {
					addr = "0.0.0.0:80"
				}
				if httpsaddr == "" {
					httpsaddr = "0.0.0.0:80"
				}
			} else {
				if addr == "" {
					addr = "127.0.0.1:8090"
				}
			}

			err := core.Serve(core.ServeConfig{
				HttpAddr:        addr,
				HttpsAddr:       httpsaddr,
				SSL:             false,
				ShowStartBanner: true,
			})

			if errors.Is(err, http.ErrServerClosed) {
				log.Println("Server closed")
				return
			}

			if err != nil {
				log.Fatalf("Server error: %v", err)
			}
		},
	}
	cmd.PersistentFlags().StringVar(
		&addr,
		"port",
		"",
		"Specify the port for the HTTP server. Defaults to 0.0.0.0:4173 unless overridden.",
	)

	return cmd
}
