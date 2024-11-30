package filehandler

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOrganizeFile(t *testing.T) {
	// Create temporary source and destination directories
	sourceDir, err := os.MkdirTemp("", "source")
	if err != nil {
		t.Fatalf("Failed to create temporary source directory: %v", err)
	}
	defer os.RemoveAll(sourceDir) // Clean up

	destDir, err := os.MkdirTemp("", "destination")
	if err != nil {
		t.Fatalf("Failed to create temporary destination directory: %v", err)
	}
	defer os.RemoveAll(destDir) // Clean up

	// Create test files in the source directory
	files := []string{"test1.txt", "test2.pdf", "test3.jpg"}
	for _, file := range files {
		filePath := filepath.Join(sourceDir, file)
		err := os.WriteFile(filePath, []byte("dummy content"), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", file, err)
		}
	}

	// Run the function
	msgChan := make(chan string)
	go func() {
		OrganizeFile(sourceDir, destDir, msgChan)
	}()

	// Capture and verify messages
	var messages []string
	for msg := range msgChan {
		messages = append(messages, msg)
	}

	// Assertions
	for _, file := range files {
		ext := strings.Split(file, ".")[1]
		destPath := filepath.Join(destDir, ext, file)
		if _, err := os.Stat(destPath); os.IsNotExist(err) {
			t.Errorf("Expected file %s to be moved to %s, but it was not", file, destPath)
		}
	}

	// Verify the final message
	if messages[len(messages)-1] != "Task completed!" {
		t.Errorf("Expected final message to be 'Task completed!', got '%s'", messages[len(messages)-1])
	}
}
