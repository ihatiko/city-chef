package cook

import (
	"city-chef/cmd/cook/domain"
	"city-chef/cmd/cook/project"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	c := &cobra.Command{
		Use:   "cook [string to print]",
		Short: "command to generate component",
	}
	c.AddCommand(domain.Command())
	c.AddCommand(project.Command())
	return c
}
