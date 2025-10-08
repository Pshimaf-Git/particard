package remove

import (
	"errors"
	"fmt"

	"github.com/Pshimaf-Git/particard/internal/storage"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var ErrNotEnougthArgs = errors.New("not enought arguments")

func NewCmd(db storage.Storage) *cobra.Command {
	removeCmd := &cobra.Command{
		Use:   "remove [id]",
		Short: "Remove a parti member by ID",
		Long: `The 'remove' command deletes a parti member record from the database
identified by their unique ID. You must provide a valid UUID as the argument.

Example:
  particard remove 123e4567-e89b-12d3-a456-426614174000`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("%w: require user id", ErrNotEnougthArgs)
			}

			id, err := uuid.Parse(args[0])
			if err != nil {
				return fmt.Errorf("member ID '%s' is not a valid UUID. Please provide a correct ID", args[0])
			}

			if err := db.RemovePartiMember(cmd.Context(), id); err != nil {
				if errors.Is(err, storage.ErrPartiMemberNotFound) {
					return fmt.Errorf("member with ID '%s' was not found", id)
				}
				return fmt.Errorf("could not remove member with ID '%s': %w", id, err)
			}

			return nil
		},
	}

	return removeCmd
}
