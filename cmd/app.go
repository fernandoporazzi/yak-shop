package cmd

import (
	"github.com/fernandoporazzi/yak-shop/app"
	"github.com/spf13/cobra"
)

var appCmd = &cobra.Command{
	Use:   "start-server",
	Short: "Start HTTP server",
	Long:  "Start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		app.Start()
	},
}
