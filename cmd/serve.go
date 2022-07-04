/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/ValerySidorin/whisper/internal/infrastructure/di"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run whisper server",
	Long:  `Run whisper server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		server, err := di.InitWebServer()
		if err != nil {
			return err
		}
		server.Run()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
