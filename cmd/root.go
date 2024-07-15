package cmd

import (
	"os"

	commentGithub "github.com/gbh-tech/github-pr-commenter/src/comments"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var GithubClient commentGithub.GithubClient

var RootCmd = &cobra.Command{
	Use:   "commenter",
	Short: "A CLI to perform operation on GitHub PR issues.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		token := os.Getenv("GITHUB_TOKEN")
		err := commentGithub.NewClient(token, &GithubClient)
		if err != nil {
			log.Fatalf("Error initializing GitHub client: %v", err)
		}
	},
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
