package cmd

import (
	"amireshoon/gow/gow"
	"fmt"

	"github.com/spf13/cobra"
)

// GowVersion current version of gow
var GowVersion = "1.0.0"

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(work)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of gow you are using.",
	Long:  `Print version number of gow you are currently using.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gow version ", gow.GetVersion())
	},
}

var work = &cobra.Command{
	Use:   "work",
	Short: "Get work by name or id",
	Long:  `Get work or todo by calling it by name or id`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD", args)
	},
}
