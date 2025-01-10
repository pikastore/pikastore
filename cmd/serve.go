package cmd

import (
	"github.com/pikastore/pikastore/api"
	"github.com/spf13/cobra"
)

func ServeCmd() *cobra.Command {
	var addr string
	cmd := &cobra.Command{
		Use:          "serve",
		Args:         cobra.ArbitraryArgs,
		Short:        "Starts the HTTP server",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && addr == "" {
				addr = "0.0.0.0:4173"
			} else if addr == "" {
				addr = "0.0.0.0:4173"
			}
			api.Serve(addr)
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
