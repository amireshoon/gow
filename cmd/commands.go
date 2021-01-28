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
	rootCmd.AddCommand(works)
	rootCmd.AddCommand(_init)
	works.AddCommand(add)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of gow you are using.",
	Long:  `Print version number of gow you are currently using.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gow version ", gow.GetVersion())
	},
}

var works = &cobra.Command{
	Use:   "works",
	Short: "Get work by name or id",
	Long:  `Get work or todo by calling it by name or id`,
	Run: func(cmd *cobra.Command, args []string) {
		gow.Parse()
	},
}

var add = &cobra.Command{
	Use:   "add [string name] [string desc]",
	Short: "Add new work",
	Long:  `Add new main work or todo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Right on way", args)
	},
}

var _init = &cobra.Command{
	Use:   "init [string project name]",
	Short: "Initialize new todo workspace",
	Long:  `Adds TODO.md if not exists and initialize new workspace`,
	Run: func(cmd *cobra.Command, args []string) {
		gow.ParseTodo(".", args...)
	},
}
