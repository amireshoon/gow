package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(work)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of gow you are using.",
	Long:  `Print version number of gow you are currently using.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD", args)
	},
}

var work = &cobra.Command{
	Use:   "work",
	Short: "Add work or get work",
	Long:  `Add new work or todo or get work by calling it by name`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD", args)
	},
}
