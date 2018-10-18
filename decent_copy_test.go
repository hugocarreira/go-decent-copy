package decentcopy

import (
	"fmt"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	execPath, _ := os.Getwd()

	fileOrigin := fmt.Sprintf(execPath + "/testCopy.txt")
	fileDestiny := fmt.Sprintf(execPath + "/folder/testCopy.txt")
	err := Copy(fileOrigin, fileDestiny)
	if err != nil {
		t.Fatalf("Error in copy file : %#v ", err.Error())
	}

}
