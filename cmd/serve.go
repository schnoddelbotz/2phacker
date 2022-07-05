/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var httpPort string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A wrapper to hugo serve",
	Long:  `Runs 2phacker's preprocessing tasks, then runs hugo serve`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		// run build
		// exec hugo serve
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVar(&httpPort, "port", ":3001", "HTTP port to serve on")
	_ = viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
}
