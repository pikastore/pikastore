package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func UpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "update",
		Args:         cobra.ArbitraryArgs,
		Short:        "Installs the latest version of pikastore",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("ts(this) aint done yet twin :skull:\n")
		},
	}
	return cmd
}
