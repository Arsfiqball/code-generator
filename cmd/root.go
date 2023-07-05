package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Templates contains access to list of templates, please provide this before executing everything
var Templates embed.FS

var rootCmd = &cobra.Command{
	Use:   "code-generator",
	Short: "Code Generator is your goto cli to generate Cynet Codec based feature structure",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("it worked!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("verbose", "V", false, "Print more detailed information")
}
