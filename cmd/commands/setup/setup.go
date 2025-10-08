package setup

import (
	"os"
	"path"

	"github.com/Pshimaf-Git/particard/internal/config"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	var db string

	setupCmd := &cobra.Command{
		Use:   "setup",
		Short: "Initialize the particard database and set up environment variables",
		Long: `The 'setup' command initializes the necessary directory structure for the particard
database and saves the database path in the configuration file.

Example:
  particard setup --db sqlite`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfgDir, err := config.GetConfigDir()
			if err != nil {
				return err
			}

			var mkdirPerm os.FileMode = 0o755

			dbDir := path.Join(cfgDir, "db", db)

			if err := os.MkdirAll(dbDir, mkdirPerm); err != nil {
				return err
			}

			return config.SaveConfig(config.Config{
				DatabaseURL: path.Join(dbDir, db+".db"),
			})
		},
	}

	setupCmd.Flags().StringVar(&db, "db", "", "database")
	setupCmd.MarkFlagRequired("db")

	return setupCmd
}
