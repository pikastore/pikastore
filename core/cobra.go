package core

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "pikastore",
		Short: "pikastore CLI",
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(serveCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of pikastore CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("pikastore version %s\n", app.version)
	},
}
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the web server",
	Run: func(cmd *cobra.Command, args []string) {
		//note: make this work
	},
}

func Execute() error {
	return rootCmd.Execute()
}
