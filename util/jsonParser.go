package util

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func ParseJSONFile[T interface{}](path string) (*T, error) {

	path, err := filepath.Abs(path)

	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var payload T
	err = json.Unmarshal(content, &payload)

	if err != nil {
		return nil, err
	}

	return &payload, nil
}
