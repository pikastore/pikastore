package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

func ServeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "serve",
		Args:         cobra.ArbitraryArgs,
		Short:        "Starts the web server (127.0.0.1:8090)",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			address := "127.0.0.1:8090"
			if len(args) > 0 {
				address = args[0]
			}
			if err := http.ListenAndServe(address, nil); err != nil {
				fmt.Println("Error starting server:", err)
			}
		},
	}
	return cmd
}
