package cmd

import (
	"log"
	"os"

	"github.com/paulcalimache/go-curriculum/internal/curriculum"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-curriculum",
	Short: "Go curriculum is a CLI tool for generating curriculum vitae",
	Long: `
Go curriculum is a CLI tool for generating curriculum vitae in pdf or html format,
based from a yaml config file.`,
	RunE: run,
}

func init() {
	rootCmd.Flags().StringP("file", "f", "", "Yaml data file")
	err := rootCmd.MarkFlagRequired("file")
	if err != nil {
		log.Fatal(err)
	}
	rootCmd.Flags().StringP("output", "o", "./output", "Output directory")
	rootCmd.Flags().StringP("template", "t", "classic", "CV Template to use")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	file, _ := cmd.Flags().GetString("file")
	output, _ := cmd.Flags().GetString("output")
	template, _ := cmd.Flags().GetString("template")

	cmd.Printf("Parsing file %s ...\n", file)
	cv, err := curriculum.ParseFile(file)
	if err != nil {
		return err
	}

	cmd.Printf("Generating curriculum vitae using '%s' template ...\n", template)
	err = cv.Render(output, template)
	if err != nil {
		return err
	}
	return nil
}
