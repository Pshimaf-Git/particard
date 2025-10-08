package get

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Pshimaf-Git/particard/internal/storage"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

const (
	defaultIndent = 2
)

var ErrUntranslatedID = errors.New("id untranslated")

func NewCmd(db storage.Storage) *cobra.Command {
	var indentSize uint
	getCmd := &cobra.Command{
		Use:   "get [id]",
		Short: "Get a parti member by ID",
		Long: `The 'get' command retrieves and displays the information of a parti member
identified by their unique ID. You must provide a valid UUID as the argument.

Example:
  particard get 123e4567-e89b-12d3-a456-426614174000`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			uid, err := uuid.Parse(args[0])
			if err != nil {
				return fmt.Errorf("member ID '%s' is not a valid UUID. Please provide a correct ID.", args[0])
			}

			member, err := db.GetPartiMember(ctx, uid)
			if err != nil {
				if errors.Is(err, storage.ErrPartiMemberNotFound) {
					return fmt.Errorf("member with ID '%s' was not found.", uid)
				}
				return fmt.Errorf("could not retrieve member with ID '%s': %w", uid, err)
			}

			indent := strings.Repeat(" ", int(indentSize))

			fmt.Printf("Member info:\n%sName: %s\n%sParti: %s\n%sRole: %s\n",
				indent, member.Name, indent, member.Parti, indent, member.Role,
			)

			return nil
		},
	}

	getCmd.Flags().UintVarP(&indentSize, "ident", "i", defaultIndent, "ident size for new lines")

	return getCmd
}
