package pikastore

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pikastore/pikastore/cmd"
	"github.com/pikastore/pikastore/core"
	"github.com/spf13/cobra"
)

type Pikastore struct {
	core.Config
	rootCmd *cobra.Command
}

func New() *Pikastore {
	return CreateAPP()
}

// Creates the App
func CreateAPP() *Pikastore {
	baseDir, _ := inspectRuntime()
	defaultData := filepath.Join(baseDir, "ps_data/")
	defaultDatabase := filepath.Join(baseDir, "ps_db/")

	config := core.Config{
		DataDir: defaultData,
		DbDir:   defaultDatabase,
	}

	ps := &Pikastore{
		Config: config,
		rootCmd: &cobra.Command{
			Use:   "pikastore",
			Short: "pikastore CLI",
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
		},
	}

	ps.rootCmd.AddCommand(cmd.ServeCmd())
	ps.rootCmd.AddCommand(cmd.UpdateCmd())

	if _, err := os.Stat(config.DataDir); os.IsNotExist(err) {
		os.Mkdir(config.DataDir, 0755)
	}
	if _, err := os.Stat(config.DbDir); os.IsNotExist(err) {
		os.Mkdir(config.DbDir, 0755)
	}

	core.Bootstrap(config)

	return ps
}

func (ps *Pikastore) Start() error {
	return ps.rootCmd.Execute()
}


func inspectRuntime() (baseDir string, withGoRun bool) {
	if strings.HasPrefix(os.Args[0], os.TempDir()) {
		withGoRun = true
		baseDir, _ = os.Getwd()
	} else {
		withGoRun = false
		baseDir = filepath.Dir(os.Args[0])
	}
	return
}
