package config

import "testing"

func TestLoad_DefaultPort(t *testing.T) {
	t.Setenv("PORT", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() returned unexpected error: %v", err)
	}

	if cfg.HTTPPort != 8080 {
		t.Errorf("HTTPPort = %d, want 8080", cfg.HTTPPort)
	}
}

func TestLoad_CustomPort(t *testing.T) {
	t.Setenv("PORT", "9090")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() returned unexpected error: %v", err)
	}

	if cfg.HTTPPort != 9090 {
		t.Errorf("HTTPPort = %d, want 9090", cfg.HTTPPort)
	}
}

func TestLoad_InvalidPort(t *testing.T) {
	t.Setenv("PORT", "hello")

	_, err := Load()
	if err == nil {
		t.Fatal("Load() expected an error for invalid PORT, got nil")
	}
}
