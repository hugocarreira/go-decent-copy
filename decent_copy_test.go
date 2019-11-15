package decentcopy

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestCopy(t *testing.T) {
	execPath, _ := os.Getwd()

	fileOrigin := fmt.Sprintf(execPath + "/testCopy.txt")
	fileDestiny := fmt.Sprintf(execPath + "/testCopied.txt")
	err := Copy(fileOrigin, fileDestiny)
	if err != nil {
		t.Fatalf("Error in copy file : %#v ", err.Error())
	}

}

func TestCopyPermissions(t *testing.T) {
	execPath, _ := os.Getwd()

	fileOrigin := fmt.Sprintf(execPath + "/testExecutable.sh")
	fileDestiny := fmt.Sprintf(execPath + "/testAlsoExecutable.sh")

	if err := Copy(fileOrigin, fileDestiny); err != nil {
		t.Fatalf("Error in copy file : %#v ", err.Error())
	}

	if err := exec.Command(fileDestiny).Run(); err != nil {
		t.Fatalf("Error in copied permissions: %#v ", err.Error())
	}
}
