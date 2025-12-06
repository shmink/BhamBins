package test

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"path/filepath"
	"testing"
)

type BinCollection struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type BinCollections struct {
	Bins []BinCollection `json:"bins"`
}

func TestBhambinsBinary(t *testing.T) {
	binPath := filepath.Join("..", "bhambins")

	cmd := exec.Command(binPath, "--postcode", "B17 0LY", "--uprn", "100070285236")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		t.Fatalf("binary failed: %v\nOutput:\n%s", err, out.String())
	}

	var data BinCollections
	if err := json.Unmarshal(out.Bytes(), &data); err != nil {
		t.Fatalf("failed to parse JSON: %v\nOutput:\n%s", err, out.String())
	}

	if len(data.Bins) == 0 {
		t.Fatalf("expected at least one bin collection, got none")
	}
	if data.Bins[0].Date == "" {
		t.Errorf("expected Date to be populated, got empty")
	}
}
