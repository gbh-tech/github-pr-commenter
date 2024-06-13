package files

import (
	"fmt"
	"os"
)

func ParseFileContent(filePath string) string {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
	}
	return string(file)
}
