/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"

	"github.com/ValerySidorin/whisper/internal/config"
	"github.com/spf13/cobra"
)

var cfgName string
var rootCmd = &cobra.Command{
	Use:   config.ProjectName,
	Short: config.ProjectName + " service",
	Long:  config.ProjectName + ` can subscribe to your GitLab webhook and send messages via Telegram.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return loadConfig(cfgName)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgName, "config", "", "config file path")
	rootCmd.AddCommand(serverCmd)
}

func loadConfig(cfgPath string) error {
	log.Println("pid: " + strconv.Itoa(os.Getpid()))
	_, err := config.LoadDefaultConfigByViper(cfgPath)
	if err != nil {
		return errors.Wrap(err, "error while loading configuration")
	}
	return err
}
