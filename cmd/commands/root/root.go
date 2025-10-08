package root

import (
	"github.com/Pshimaf-Git/particard/cmd/commands/get"
	"github.com/Pshimaf-Git/particard/cmd/commands/new"
	"github.com/Pshimaf-Git/particard/cmd/commands/remove"
	"github.com/Pshimaf-Git/particard/cmd/commands/setup"
	"github.com/Pshimaf-Git/particard/cmd/commands/update"
	"github.com/Pshimaf-Git/particard/internal/storage"

	"github.com/spf13/cobra"
)

func NewRootCmd(storage storage.Storage) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "particard [command] [args]",
		Short: "A CLI tool for managing parti members",
		Long: `particard is a command-line interface (CLI) tool designed to manage parti members.
It allows you to create, retrieve, update, and remove member records efficiently.`,
	}

	rootCmd.AddCommand(get.NewCmd(storage))
	rootCmd.AddCommand(new.NewCmd(storage))
	rootCmd.AddCommand(remove.NewCmd(storage))
	rootCmd.AddCommand(update.NewCmd(storage))
	rootCmd.AddCommand(setup.NewCmd())

	return rootCmd
}
