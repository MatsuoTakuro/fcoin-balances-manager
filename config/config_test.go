package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	wantPort := 3333
	t.Setenv("HTTP_PORT", fmt.Sprint(wantPort))

	got, err := New()
	if err != nil {
		t.Fatalf("failed to create config: %v", err)
	}
	if got.HttpPort != wantPort {
		t.Errorf("want %d, but %d", wantPort, got.HttpPort)
	}

	// default値の検証
	wantDBPort := 3306
	if got.DBPort != wantDBPort {
		t.Errorf("want %d, but %d", wantDBPort, got.DBPort)
	}

}
