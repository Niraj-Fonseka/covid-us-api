package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filename, path string) ([]byte, error) {

	fullPath := buildFilePath(filename, path)
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	return ioutil.ReadAll(file)
}

func SaveFile(filename, path string, data []byte) error {
	fullPath := buildFilePath(filename, path)

	err := ioutil.WriteFile(fullPath, data, 0644)

	return err
}

func buildFilePath(filename, path string) string {
	if path == "" {
		return filename
	}
	return fmt.Sprintf("%s/%s", filename, path)
}
