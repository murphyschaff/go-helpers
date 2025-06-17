package helpers

import (
	"fmt"
	"io"
	"os"
)

// CopyFile copies contents from source to dest using io.CopyBuffer
func CopyFile(source string, dest string) error {
	src, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("unable to open file %s: %s", source, err)
	}
	dst, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("unable to create destination file %s: %s", dest, err)
	}
	defer src.Close()
	defer dst.Close()

	buf := make([]byte, 32*1024)
	_, err = io.CopyBuffer(dst, src, buf)
	if err != nil {
		return fmt.Errorf("error copying file: %s", err)
	}
	return nil
}
