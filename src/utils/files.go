package files

import (
	"fmt"
	"os"
	"unicode/utf8"
)

const MaxCommentBodyLength = 65000

const truncationNotice = "\n\n---\n\n_Comment truncated to 65,000 characters._"

func TruncateCommentBody(body string) string {
	if utf8.RuneCountInString(body) <= MaxCommentBodyLength {
		return body
	}

	maxContentRunes := MaxCommentBodyLength - utf8.RuneCountInString(truncationNotice)
	if maxContentRunes < 0 {
		maxContentRunes = 0
	}

	return string([]rune(body)[:maxContentRunes]) + truncationNotice
}

func ParseFileContent(filePath string) string {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
	}
	return string(file)
}

func GetCommentBody(content, filePath string) string {
	if filePath != "" {
		return TruncateCommentBody(ParseFileContent(filePath))
	}
	return TruncateCommentBody(content)
}
