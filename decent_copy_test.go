package decentcopy

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestCopy(t *testing.T) {
	execPath, _ := os.Getwd()

	fileOrigin := filepath.Join(execPath, "testCopy.txt")
	fileDestiny := filepath.Join(execPath, "testCopied.txt")
	err := Copy(fileOrigin, fileDestiny)
	if err != nil {
		t.Fatalf("Error in copy file : %#v ", err.Error())
	}

}

func TestCopyPermissions(t *testing.T) {
	execPath, _ := os.Getwd()

	fileOrigin := filepath.Join(execPath, "testExecutable.sh")
	fileDestiny := filepath.Join(execPath, "testAlsoExecutable.sh")

	if err := Copy(fileOrigin, fileDestiny); err != nil {
		t.Fatalf("Error in copy file : %#v ", err.Error())
	}

	if err := exec.Command(fileDestiny).Run(); err != nil {
		t.Fatalf("Error in copied permissions: %#v ", err.Error())
	}
}

func TestOriginDoesNotExist(t *testing.T) {
	execPath, _ := os.Getwd()

	fileOrigin := filepath.Join(execPath, "notThere.txt")
	fileDestiny := filepath.Join(execPath, "neverThere.txt")

	if err := Copy(fileOrigin, fileDestiny); err == nil {
		t.Fatalf("Should have returned an error because %v does not exist", fileOrigin)
	}
}

func TestPermissionsUnchangable(t *testing.T) {
	execPath, _ := os.Getwd()

	fileOrigin := filepath.Join(execPath, "testHammertime.sh")
	fileDestiny := filepath.Join(execPath, "maybeRun.sh")

	if err := Copy(fileOrigin, fileDestiny); err == nil {
		t.Fatalf("Should have returned an error because %v does not exist", fileOrigin)
	}
}
