// Package decentcopy copy files in simple form
package decentcopy

import (
	"io"
	"os"
)

// Copy use to copy files
func Copy(filepathOrigin, filepathDestiny string) error {
	srcFile, err := os.Open(filepathOrigin)
	defer srcFile.Close()

	destFile, err := os.Create(filepathDestiny)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)

	err = destFile.Sync()

	if err != nil {
		return err
	}
	return nil
}
