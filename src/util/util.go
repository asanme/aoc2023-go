package util

import (
	"os"
)

func ReadFileBytes(filename string) ([]byte, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path := workingDir + "/src/data/" + filename

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}
