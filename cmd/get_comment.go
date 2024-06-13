package cmd

import (
	commentGithub "github.com/gbh-tech/github-pr-commenter/src"

	"github.com/spf13/cobra"
)

var getCommentPrCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get message ID based on text",
	Example: "atlas comment get [flags]",
	Run: func(cmd *cobra.Command, args []string) {
		pull, _ := cmd.Flags().GetInt("pull")
		repo, _ := cmd.Flags().GetString("repo")
		org, _ := cmd.Flags().GetString("org")
		content, _ := cmd.Flags().GetString("content")
		filePath, _ := cmd.Flags().GetString("filePath")

		commentGithub.GetUserComments(pull, org, repo, content, filePath)
	},
}

func init() {
	RootCmd.AddCommand(getCommentPrCmd)

	getCommentPrCmd.PersistentFlags().StringP("content", "c", "", "Comment Content")
	getCommentPrCmd.PersistentFlags().StringP("filePath", "f", "", "File path")
}
