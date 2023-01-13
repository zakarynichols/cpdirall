package dir

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Cp(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, strings.TrimPrefix(path, src))

		if info.IsDir() {
			err := os.MkdirAll(dstPath, info.Mode())
			if err != nil {
				return err
			}
		} else {
			_, err := cpFile(path, dstPath) // return the written bytes?
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func cpFile(src, dst string) (int64, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()
	written, err := io.Copy(dstFile, srcFile)
	return written, err
}
