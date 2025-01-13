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

type Config struct {
	Data  string
	DbCon string
}

// Creates the App
func CreateAPP() *Pikastore {
	baseDir, _ := inspectRuntime()
	var DefaultData = filepath.Join(baseDir, "ps_data")
	var DefaultDatabase = filepath.Join(baseDir, "ps_db")
	ps := &Pikastore{
		rootCmd: &cobra.Command{
			Use:   "pikastore",
			Short: "pikastore CLI",
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
		},
	}
	core.NewApp(core.Config{
		Data:  DefaultData,
		DbCon: DefaultDatabase,
	})
	ps.rootCmd.AddCommand(cmd.ServeCmd())
	ps.rootCmd.AddCommand(cmd.UpdateCmd())

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
