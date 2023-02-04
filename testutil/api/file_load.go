package api

import (
	"os"
	"testing"
)

func LoadJSONFile(t *testing.T, path string) []byte {
	t.Helper()

	json, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read file from %q: %v", path, err)
	}

	return json
}
