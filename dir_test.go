package dir

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCp(t *testing.T) {
	// Test copying a single file
	file, err := os.Create("example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	_, err = file.Write([]byte("example file content"))
	if err != nil {
		t.Fatal(err)
	}

	dest := filepath.Join(os.TempDir(), "example_copy")
	defer os.RemoveAll(dest)

	if err := Cp(file.Name(), dest); err != nil {
		t.Fatal(err)
	}

	copiedFile, err := os.ReadFile(dest)
	if err != nil {
		t.Fatal(err)
	}

	if string(copiedFile) != "example file content" {
		t.Error("Copied file content does not match the original")
	}

	// Test copying a directory
	srcDir, err := os.MkdirTemp("", "src")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(srcDir)

	file1, err := os.Create(filepath.Join(srcDir, "file1"))
	if err != nil {
		t.Fatal(err)
	}
	defer file1.Close()

	file2, err := os.Create(filepath.Join(srcDir, "file2"))
	if err != nil {
		t.Fatal(err)
	}
	defer file2.Close()

	err = os.Mkdir(filepath.Join(srcDir, "subdir"), 0755)
	if err != nil {
		t.Fatal(err)
	}

	subdir := filepath.Join(srcDir, "subdir")

	subfile, err := os.Create(filepath.Join(subdir, "subfile"))
	if err != nil {
		t.Fatal(err)
	}
	defer subfile.Close()

	destDir := filepath.Join(os.TempDir(), "dest")
	defer os.RemoveAll(destDir)

	if err := Cp(srcDir, destDir); err != nil {
		t.Fatal(err)
	}
	// check if the copied directory is identical to the source directory
	if !dirsIdentical(srcDir, destDir) {
		t.Error("Copied directory does not match the original")
	}
}

func dirsIdentical(dir1, dir2 string) bool {
	files1, _ := os.ReadDir(dir1)
	files2, _ := os.ReadDir(dir2)

	if len(files1) != len(files2) {
		return false
	}

	for i, file1 := range files1 {
		file2 := files2[i]
		if file1.Name() != file2.Name() || file1.IsDir() != file2.IsDir() {
			return false
		}
		if file1.IsDir() {
			if !dirsIdentical(filepath.Join(dir1, file1.Name()), filepath.Join(dir2, file2.Name())) {
				return false
			}
		} else {
			bytes1, _ := os.ReadFile(filepath.Join(dir1, file1.Name()))
			bytes2, _ := os.ReadFile(filepath.Join(dir2, file2.Name()))
			if string(bytes1) != string(bytes2) {
				return false
			}
		}
	}
	return true
}
