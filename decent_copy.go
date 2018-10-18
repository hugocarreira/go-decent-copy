// Package decentcopy copy files in simple form
package decentcopy

import (
	"io"
	"os"
)

// Copy use to copy files
func Copy(filepathOrigin, filepathDestiny string) error {
	srcFile, err := os.Open(filepathOrigin)

	destFile, err := os.Create(filepathDestiny)

	_, err = io.Copy(destFile, srcFile)

	err = destFile.Sync()

	return err
}
