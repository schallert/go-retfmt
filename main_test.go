package main

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	rootDir = "."
	badFiles = []string{}

	err := filepath.Walk(rootDir, fileCheck)
	if err != nil {
		t.Errorf("Error walking files: %s\n", err.Error())
	}

	t.Log("badFiles = ", badFiles)
	if len(badFiles) != 2 {
		t.Error("Error: len(badFiles) should be 2 but is %d\n", len(badFiles))
	}

	if !reflect.DeepEqual(badFiles, []string{"test/bad1.go", "test/bad2.go"}) {
		t.Error("Error: badfiles should be [test/bad1.go, test/bad2.go] but is ", badFiles)
	}
}

func TestRetVal(t *testing.T) {
	rootDir = "."
	badFiles = []string{}

	err := filepath.Walk(rootDir, fileCheck)
	if err != nil {
		t.Errorf("Error walking files: %s\n", err.Error())
	}

	if ret != 2 {
		t.Errorf("Error: ret should be 2 but is %d\n", ret)
	}

	rootDir = "foobar"
	err = filepath.Walk(rootDir, fileCheck)
	if err == nil {
		t.Error("Should have encountered error walking non-existent directory")
	}

	if ret != 1 {
		t.Errorf("Error: ret should be 1 but is %d\n", ret)
	}
}
