package main

import (
	"context"
	"log"
	"os"

	"github.com/Pshimaf-Git/particard/cmd/commands/root"
	"github.com/Pshimaf-Git/particard/cmd/commands/setup"
	"github.com/Pshimaf-Git/particard/internal/config"
	"github.com/Pshimaf-Git/particard/internal/storage/sqlite"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rootCmd := root.NewRootCmd(nil) // Pass nil for storage initially

	// Check if the setup command is being run
	if len(os.Args) > 1 && os.Args[1] == setup.NewCmd().Use {
		setupCmd := setup.NewCmd()
		rootCmd.AddCommand(setupCmd)
		if err := rootCmd.Execute(); err != nil {
			log.Fatal(err)
		}
		return
	}

	// For other commands, load config and initialize storage
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlite.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Re-create root command with actual storage
	rootCmd = root.NewRootCmd(db)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
