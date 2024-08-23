package util

import (
	"encoding/json"
	"os"
)

func ParseDatabaseCredentials(path string) (DBCredential, error) {
	content, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var payload DBCredential

	err = json.Unmarshal(content, &payload)

	if err != nil {
		return nil, err
	}

	return payload, nil
}
