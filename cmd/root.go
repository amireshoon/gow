package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	// Work flags
	fName string
	fDate string

	rootCmd = &cobra.Command{
		Use:   "gow",
		Short: "A work and todo cli tool management",
		Long: `Gow is a tool for managing adding and showing works and todo
		without leaving terminal.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gow.yaml)")
	rootCmd.PersistentFlags().StringVarP(&fName, "name", "n", "", "name flag for works and more")
	rootCmd.PersistentFlags().StringVarP(&fDate, "date", "d", "null", "date flag for works and more")
	viper.SetDefault("date", nil)
	viper.SetConfigName("gow")
	viper.AddConfigPath("/etc/gow/")
	viper.AddConfigPath("$HOME/.gow")
	viper.AddConfigPath(".")
	viper.Set("gow_path", "$HOME/work.go")
	// rootCmd.AddCommand(addCmd)
	// rootCmd.AddCommand(initCmd)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		// Search config in home directory with name ".gow" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gow")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
