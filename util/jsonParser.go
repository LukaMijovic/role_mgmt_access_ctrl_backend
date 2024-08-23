package util

import (
	"encoding/json"
	"os"
	"path/filepath"

	creds "lukamijovic.com/role-mgmt-access-ctrl/credentials"
)

func ParseDatabaseCredentials(path string) (*creds.DBCredential, error) {

	path, err := filepath.Abs(path)

	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var payload creds.DBCredential

	err = json.Unmarshal(content, &payload)

	if err != nil {
		return nil, err
	}

	return &payload, nil
}
