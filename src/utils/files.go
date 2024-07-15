package files

import (
	"fmt"
	"os"
)

// Parse content from a file to string
func ParseFileContent(filePath string) string {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
	}
	return string(file)
}

// getCommentBody returns the comment body from either content or filePath
func GetCommentBody(content, filePath string) string {
	if filePath != "" {
		return ParseFileContent(filePath)
	}
	return content
}
