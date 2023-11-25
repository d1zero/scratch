/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/d1zero/scratch/internal"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generates new service",
	Long:  `Generates new service`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.New(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.PersistentFlags().Bool("postgres", false, "adds postgres support to service")
}
