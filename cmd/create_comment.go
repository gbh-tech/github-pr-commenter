package cmd

import (
	commentGithub "github.com/gbh-tech/github-pr-commenter/internal"

	"github.com/spf13/cobra"
)

var createCommentPrCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new"},
	Short:   "Create new comment on PR",
	Example: "atlas comment create [flags]",
	Run: func(cmd *cobra.Command, args []string) {
		pull, _ := cmd.Flags().GetInt("pull")
		repo, _ := cmd.Flags().GetString("repo")
		org, _ := cmd.Flags().GetString("org")
		content, _ := cmd.Flags().GetString("content")
		filePath, _ := cmd.Flags().GetString("filePath")

		commentGithub.CreateComment(pull, org, repo, content, filePath)
	},
}

func init() {
	RootCmd.AddCommand(createCommentPrCmd)

	createCommentPrCmd.PersistentFlags().StringP("content", "c", "", "Comment Content")
	createCommentPrCmd.PersistentFlags().StringP("filePath", "f", "", "File path")
}
