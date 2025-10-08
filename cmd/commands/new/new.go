package new

import (
	"fmt"

	"github.com/Pshimaf-Git/particard/internal/models"
	"github.com/Pshimaf-Git/particard/internal/storage"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	defaultParti = "German Legion"
	defaultRole  = "member"
)

func NewCmd(db storage.Storage) *cobra.Command {
	var (
		parti string
		name  string
		role  string
	)

	newCmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new parti member",
		Long: `The 'new' command creates a new parti member record in the database.
You can specify the parti, name, and role using flags. If not provided,
default values will be used for parti ("German Legion") and role ("member").

Example:
  particard new --name "John Doe" --parti "New Alliance" --role "Recruit"
  particard new -n "Jane Smith"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			member := models.NewPartiMember(uuid.New(), parti, name, role)

			uid, err := db.CreatePartiMember(ctx, member)
			if err != nil {
				return fmt.Errorf("could not create new member: %w", err)
			}

			fmt.Println("uid:", uid)

			return nil
		},
	}

	newCmd.Flags().StringVarP(&parti, "parti", "p", defaultParti, "parti name")
	newCmd.Flags().StringVarP(&name, "name", "n", "", "name of parti member")
	newCmd.Flags().StringVarP(&role, "role", "r", defaultRole, "role of parti member")

	return newCmd
}
