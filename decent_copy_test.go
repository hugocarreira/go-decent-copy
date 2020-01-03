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

func TestSourceFileExistsWithPermissionDenied(t *testing.T) {
	execPath, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error while getting the working directory %v", err.Error())
	}

	fileOrigin := filepath.Join(execPath, "testExecutable.sh")

	err = os.Chmod(fileOrigin, 0000)
	if err != nil {
		t.Fatalf("Error while changing the permission to 0000: %v", err.Error())
	}

	defer func() {
		err = os.Chmod(fileOrigin, 0775)
		if err != nil {
			t.Fatalf("Error while changing the permission to 0775: %v", err.Error())
		}
	}()

	fileDestiny := filepath.Join(execPath, "neverThere.txt")

	if err := Copy(fileOrigin, fileDestiny); err == nil {
		t.Fatalf("Should have returned an error because the user cannot open the file")
	}
}

func TestDestFileExistsWithPermissionDenied(t *testing.T) {
	execPath, _ := os.Getwd()

	fileOrigin := filepath.Join(execPath, "testExecutable.sh")
	fileDestiny := filepath.Join("/dev/null/wrong-path")

	if err := Copy(fileOrigin, fileDestiny); err == nil {
		t.Fatalf("Should have returned an error because %v does not exist", fileDestiny)
	}
}
