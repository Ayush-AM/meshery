package design

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetUniqueFilename(t *testing.T) {
	tempDir := t.TempDir()

	t.Run("no existing file returns original filename", func(t *testing.T) {
		filename := filepath.Join(tempDir, "pattern_1234.yaml")
		got := getUniqueFilename(filename)
		if got != filename {
			t.Errorf("expected %s, got %s", filename, got)
		}
	})

	t.Run("single collision appends (1) before extension", func(t *testing.T) {
		filename := filepath.Join(tempDir, "design_single.yaml")
		if err := os.WriteFile(filename, []byte("data"), 0o644); err != nil {
			t.Fatalf("failed to create dummy file: %v", err)
		}

		expected := filepath.Join(tempDir, "design_single(1).yaml")
		got := getUniqueFilename(filename)
		if got != expected {
			t.Errorf("expected %s, got %s", expected, got)
		}
	})

	t.Run("multiple collisions increment index before extension", func(t *testing.T) {
		filename := filepath.Join(tempDir, "design_multi.yaml")
		file1 := filepath.Join(tempDir, "design_multi(1).yaml")
		file2 := filepath.Join(tempDir, "design_multi(2).yaml")

		for _, f := range []string{filename, file1, file2} {
			if err := os.WriteFile(f, []byte("data"), 0o644); err != nil {
				t.Fatalf("failed to create dummy file: %v", err)
			}
		}

		expected := filepath.Join(tempDir, "design_multi(3).yaml")
		got := getUniqueFilename(filename)
		if got != expected {
			t.Errorf("expected %s, got %s", expected, got)
		}
	})

	t.Run("preserves multi-part extension .tar.gz", func(t *testing.T) {
		filename := filepath.Join(tempDir, "design_archive.tar.gz")
		if err := os.WriteFile(filename, []byte("data"), 0o644); err != nil {
			t.Fatalf("failed to create dummy file: %v", err)
		}

		expected := filepath.Join(tempDir, "design_archive(1).tar.gz")
		got := getUniqueFilename(filename)
		if got != expected {
			t.Errorf("expected %s, got %s", expected, got)
		}
	})

	t.Run("handles file with no extension", func(t *testing.T) {
		filename := filepath.Join(tempDir, "design_noext")
		if err := os.WriteFile(filename, []byte("data"), 0o644); err != nil {
			t.Fatalf("failed to create dummy file: %v", err)
		}

		expected := filepath.Join(tempDir, "design_noext(1)")
		got := getUniqueFilename(filename)
		if got != expected {
			t.Errorf("expected %s, got %s", expected, got)
		}
	})
}
