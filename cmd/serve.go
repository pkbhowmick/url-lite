package cmd

import (
	"github.com/pkbhowmick/url-lite/apis"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Long:  "Serve the url-lite server",
	Short: "Server the url-lite server",
	Run: func(cmd *cobra.Command, args []string) {
		apis.Serve()
	},
}
