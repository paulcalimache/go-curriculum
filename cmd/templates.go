package cmd

import (
	"github.com/paulcalimache/go-curriculum/internal/templates"
	"github.com/spf13/cobra"
)

var templatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "List availables templates",
	Long:  `List all curriculum vitae templates availables`,
	RunE:  listTemplates,
}

func init() {
	rootCmd.AddCommand(templatesCmd)
}

func listTemplates(cmd *cobra.Command, args []string) error {
	rootCmd.Print(templates.GetTemplatesList())
	return nil
}
