package githubComments

import (
	"context"
	"fmt"
	"strings"

	files "github.com/gbh-tech/github-pr-commenter/src/utils"

	"github.com/charmbracelet/log"
	"github.com/google/go-github/v61/github"
)

type GithubClient struct {
	Ctx    context.Context
	Client *github.Client
}

// NewClient creates a new GithubClient with the provided token
func NewClient(token string, client *GithubClient) error {
	if token == "" {
		return fmt.Errorf("Unauthorized: GitHub Token not present.")
	}
	client.Ctx = context.Background()
	client.Client = github.NewClient(nil).WithAuthToken(token)
	return nil
}

func (client *GithubClient) GetUserComments(pull int, org, repo, content, filePath string) (int64, error) {
	comments, _, err := client.Client.Issues.ListComments(client.Ctx, org, repo, pull, nil)
	if err != nil {
		return 0, fmt.Errorf("PR ID: %v\nError listing comments: %v", pull, err)
	}

	commentBody := files.GetCommentBody(content, filePath)
	for _, comment := range comments {
		if strings.Contains(*comment.Body, commentBody) {
			return *comment.ID, nil
		}
	}

	return 0, fmt.Errorf("PR ID: %v\nComment not found", pull)
}

func (client *GithubClient) CreateComment(pull int, org, repo, content, filePath string) (*github.IssueComment, error) {
	commentBody := files.GetCommentBody(content, filePath)
	comment := &github.IssueComment{Body: github.String(commentBody)}

	commentResp, _, err := client.Client.Issues.CreateComment(client.Ctx, org, repo, pull, comment)
	if err != nil {
		return nil, fmt.Errorf("PR ID: %v\nError creating comment: %v", pull, err)
	}
	log.Info("Comment Created successfully")

	return commentResp, nil
}

func (client *GithubClient) UpdateComment(commentID int64, org, repo, content, filePath string) (*github.IssueComment, error) {
	commentBody := files.GetCommentBody(content, filePath)
	comment := &github.IssueComment{Body: github.String(commentBody)}

	commentResp, _, err := client.Client.Issues.EditComment(client.Ctx, org, repo, commentID, comment)
	if err != nil {
		return nil, fmt.Errorf("Comment ID: %v\nError updating comment: %v", commentID, err)
	}
	log.Info("Comment updated successfully")

	return commentResp, nil
}
