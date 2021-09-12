package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "yakshop",
		Short: "A cli to work with YakShop",
		Long:  "YakShop CLI allows you to query data or start an HTTP server",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	getDataCmd.PersistentFlags().StringP("file", "f", "", "XML file to read from")
	getDataCmd.PersistentFlags().Int32P("days", "d", 0, "Elapsed time in days")
	getDataCmd.MarkPersistentFlagRequired("file")
	getDataCmd.MarkPersistentFlagRequired("days")

	rootCmd.AddCommand(getDataCmd)
	rootCmd.AddCommand(appCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
