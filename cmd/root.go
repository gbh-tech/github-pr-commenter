package cmd

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "commenter",
	Short: "A CLI to perform operation on GitHub PR issues.",
}

func init() {
	RootCmd.PersistentFlags().StringP("repo", "r", "", "Target repository")
	RootCmd.PersistentFlags().IntP("pull", "p", 0, "Target PR")
	RootCmd.PersistentFlags().StringP("org", "o", "", "GitHub Organization name")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
