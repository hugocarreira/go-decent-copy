/*

decentcopy provides a copy files for humans

Usage

	package main

	import "github.com/hugocarreira/go-decent-copy"

	func main() {
		execPath, _ := os.Getwd()

		fileOrigin := "/testCopy.txt"
		fileDestiny := fmt.Sprintf(execPath + "/folder/" + fileOrigin)
		err := Copy(fileOrigin, fileDestiny)
		if err != nil {
			fmt.Printf("Error in copy file : %#v ", err.Error())
		}
	}

*/

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

	return err
}
