// Package decentcopy copy files in simple form
package decentcopy

import (
	"io"
	"os"
)

// ExecOnErrorWhileDefer is executed when the file.Close() function returns an
// error while closing the source or destination files.
var ExecOnErrorWhileDefer func(error)

// Copy use to copy files
func Copy(filepathOrigin, filepathDestiny string) error {
	srcInfo, err := os.Stat(filepathOrigin)
	if err != nil {
		return err
	}

	srcFile, err := os.Open(filepathOrigin)
	if err != nil {
		return err
	}
	defer func() {
		err := srcFile.Close()
		if ExecOnErrorWhileDefer != nil {
			ExecOnErrorWhileDefer(err)
		}
	}()

	destFile, err := os.Create(filepathDestiny)
	if err != nil {
		return err
	}
	defer func() {
		err := destFile.Close()
		if ExecOnErrorWhileDefer != nil {
			ExecOnErrorWhileDefer(err)
		}
	}()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return err
	}

	if err := os.Chmod(filepathDestiny, srcInfo.Mode()); err != nil {
		return err
	}

	return destFile.Sync()
}
