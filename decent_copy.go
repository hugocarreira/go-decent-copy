// Package decentcopy copy files in simple form
package decentcopy

import (
	"io"
	"os"
)

// Copy use to copy files
func Copy(filepathOrigin, filepathDestiny string) error {
	srcInfo, err := os.Stat(filepathOrigin)
	if err != nil {
		return err
	}

	srcFile, _ := os.Open(filepathOrigin)
	defer srcFile.Close()

	destFile, _ := os.Create(filepathDestiny)
	defer destFile.Close()

	io.Copy(destFile, srcFile)

	if err := os.Chmod(filepathDestiny, srcInfo.Mode()); err != nil {
		return err
	}

	return destFile.Sync()
}
