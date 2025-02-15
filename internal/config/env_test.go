package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Create a temporary .env file for testing
	content := []byte("APP_NAME=tinystack\nPORT=8081\n")
	tmpFile, err := os.CreateTemp("", "test.env")
	if err != nil {
		t.Fatal("Failed to create temp file:", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up after test

	if _, err := tmpFile.Write(content); err != nil {
		t.Fatal("Failed to write to temp file:", err)
	}

	// Seek to start of file for reading
	if _, err := tmpFile.Seek(0, 0); err != nil {
		t.Fatal("Failed to seek file:", err)
	}

	// Test the Load function
	_, err = Load(tmpFile)
	if err != nil {
		t.Fatal("Failed to load env file:", err)
	}

	// Test the environment variables
	expectedAppName := "tinystack"
	if got := os.Getenv("APP_NAME"); got != expectedAppName {
		t.Errorf("Expected APP_NAME to be %s, got %s", expectedAppName, got)
	}

	expectedPort := "8081"
	if got := os.Getenv("PORT"); got != expectedPort {
		t.Errorf("Expected PORT to be %s, got %s", expectedPort, got)
	}
}
