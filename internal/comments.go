package githubComments

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/gbh-tech/github-pr-commenter/utils/files"

	"github.com/charmbracelet/log"
	"github.com/google/go-github/v61/github"
)

type GithubClient struct {
	Ctx    context.Context
	Client *github.Client
}

func DeclareClient(client *GithubClient) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	client.Ctx = context.Background()
	client.Client = github.NewClient(nil).WithAuthToken(token)
}

func GetUserComments(pull int, org, repo, content, filePath string) {
	var client GithubClient
	DeclareClient(&client)

	comments, _, err := client.Client.Issues.ListComments(client.Ctx, org, repo, pull, nil)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	commentBody := content
	if filePath != "" {
		commentBody = files.ParseFileContent(filePath)
	}
	for _, comment := range comments {
		if strings.Contains(*comment.Body, commentBody) {
			fmt.Println(*comment.ID)
			return
		}
	}
	log.Error("Comment not found!")
}

func CreateComment(pull int, org, repo, content, filePath string) {
	var client GithubClient
	DeclareClient(&client)

	commentBody := content
	if filePath != "" {
		commentBody = "```\n" + files.ParseFileContent(filePath) + "\n```"
	}

	comment := &github.IssueComment{
		Body: github.String(commentBody),
	}
	commentResp, _, err := client.Client.Issues.CreateComment(client.Ctx, org, repo, pull, comment)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	fmt.Printf("Comment was successfully added!\nContent: %v\n", *commentResp.Body)
}

func UpdateComment(commentID int64, org, repo, content, filePath string) {
	var client GithubClient
	DeclareClient(&client)

	commentBody := "```\n" + content + "```"
	if filePath != "" {
		commentBody = "```\n" + files.ParseFileContent(filePath) + "```"
	}

	comment := &github.IssueComment{
		Body: github.String(commentBody),
	}
	commentResp, _, err := client.Client.Issues.EditComment(client.Ctx, org, repo, commentID, comment)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	fmt.Printf("Comment was successfully Updated!\nContent: %v\n", *commentResp.Body)
}
