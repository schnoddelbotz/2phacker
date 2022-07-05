/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var AppVersion string
var UserHomeDir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "2phacker",
	Short: "",
	Long:  ``,
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

	rootCmd.PersistentFlags().StringP("docs-dir", "d", homePath("TUBE_CONTENT/docs"), "HUGO repository checkout to work on")
	rootCmd.PersistentFlags().StringP("movies-dir", "m", homePath("TUBE_CONTENT/movies"), "top-level folder with videos (.mp4, .mov)")
	rootCmd.PersistentFlags().StringP("sounds-dir", "s", homePath("TUBE_CONTENT/sounds"), "top-level folder with music (.mp3)")
	rootCmd.PersistentFlags().StringP("images-dir", "i", homePath("TUBE_CONTENT/images"), "top-level folder with images (.jpg)")
	rootCmd.PersistentFlags().StringP("webroot", "w", homePath("TUBE_WEBROOT"), "Local HUGO output directory") // ? S3 rsync hardlink ...
	rootCmd.PersistentFlags().StringP("templates", "t", "", "Path to custom templates for movies, sounds and images; CSS etc.")

	_ = viper.BindPFlag("webroot", rootCmd.PersistentFlags().Lookup("webroot"))

	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("2PH")
}

func homePath(dir string) string {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return filepath.Join(homeDir, dir)
}

func initConfig() {
	viper.AutomaticEnv()
}
