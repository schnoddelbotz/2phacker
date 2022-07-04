/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var AppVersion string
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "2phacker",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.2phacker.yaml)")

	rootCmd.PersistentFlags().StringP("docs-dir", "d", "~/TUBE_CONTENT/docs", "HUGO repository checkout to work on")
	rootCmd.PersistentFlags().StringP("movies-dir", "m", "~/TUBE_CONTENT/movies", "top-level folder with videos (.mp4, .mov)")
	rootCmd.PersistentFlags().StringP("sounds-dir", "s", "~/TUBE_CONTENT/sounds", "top-level folder with music (.mp3)")
	rootCmd.PersistentFlags().StringP("images-dir", "i", "~/TUBE_CONTENT/images", "top-level folder with images (.jpg)")
	rootCmd.PersistentFlags().StringP("webroot", "w", "~/TUBE_WEBROOT/", "Local HUGO output directory") // ? S3 rsync hardlink ...

	_ = viper.BindPFlag("webroot", rootCmd.PersistentFlags().Lookup("webroot"))

	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("2PH")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".2phacker" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".2phacker")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
