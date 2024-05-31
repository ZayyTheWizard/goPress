package EnFile

import (
	"os"
)

func FetchFileContent(FileName string) (string, error) {
	File, err := os.ReadFile(FileName)
	if err != nil {
		return "", err
	}

	Content := string(File[:])

	return Content, err
}
