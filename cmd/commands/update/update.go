package update

import (
	"context"
	"fmt"

	"github.com/Pshimaf-Git/particard/internal/models"
	"github.com/Pshimaf-Git/particard/internal/storage"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func NewCmd(db storage.Storage) *cobra.Command {
	var parti, name, role string

	updateCmd := &cobra.Command{
		Use:   "update [id]",
		Short: "Update information of a parti member",
		Long: `The 'update' command modifies the information of an existing parti member.
You must provide the member's unique ID as an argument, and then use flags
to specify the fields you wish to update (parti, name, or role).

Example:
  particard update 123e4567-e89b-12d3-a456-426614174000 --name "New Name" -p "New Parti"
  particard update 123e4567-e89b-12d3-a456-426614174000 -r "New Role"`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			memberID, err := uuid.Parse(args[0])
			if err != nil {
				return fmt.Errorf("member ID '%s' is not a valid UUID. Please provide a correct ID", args[0])
			}

			member := &models.PartiMember{
				ID:    memberID,
				Parti: parti,
				Name:  name,
				Role:  role,
			}

			err = db.UpdatePartiMember(context.Background(), memberID, member)
			if err != nil {
				if err == storage.ErrPartiMemberNotFound {
					return fmt.Errorf("member with ID '%s' was not found", memberID)
				}
				return fmt.Errorf("could not update member with ID '%s': %w", memberID, err)
			}

			fmt.Printf("Member %s updated successfully.\n", memberID)
			return nil
		},
	}

	updateCmd.Flags().StringVarP(&parti, "parti", "p", "", "New parti of the member")
	updateCmd.Flags().StringVarP(&name, "name", "n", "", "New name of the member")
	updateCmd.Flags().StringVarP(&role, "role", "r", "", "New role of the member")

	return updateCmd
}
