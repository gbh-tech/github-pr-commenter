package cmd

import (
	commentGithub "github.com/gbh-tech/github-pr-commenter/src"

	"github.com/spf13/cobra"
)

var updateCommentPrCmd = &cobra.Command{
	Use:     "update",
	Short:   "Update comment on PR",
	Example: "atlas comment update [flags]",
	Run: func(cmd *cobra.Command, args []string) {
		repo, _ := cmd.Flags().GetString("repo")
		org, _ := cmd.Flags().GetString("org")
		content, _ := cmd.Flags().GetString("content")
		filePath, _ := cmd.Flags().GetString("filePath")
		commentID, _ := cmd.Flags().GetInt64("commentID")

		commentGithub.UpdateComment(commentID, org, repo, content, filePath)
	},
}

func init() {
	RootCmd.AddCommand(updateCommentPrCmd)

	updateCommentPrCmd.PersistentFlags().StringP("content", "c", "", "Comment Content")
	updateCommentPrCmd.PersistentFlags().StringP("filePath", "f", "", "File path")
	updateCommentPrCmd.PersistentFlags().Int64P("commentID", "i", 0, "Comment ID")

}
