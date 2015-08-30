package main

import (
	"fmt"
	"path/filepath"
	"reflect"
	"testing"
)

func setup() {
	rootDir = "."
	badFiles = []string{}
	ret = 0
}

func TestFunc(t *testing.T) {
	setup()

	err := filepath.Walk(rootDir, fileCheck)
	if err != nil {
		t.Errorf("Error walking files: %s\n", err.Error())
	}

	t.Log("badFiles = ", badFiles)
	if len(badFiles) != 2 {
		t.Errorf("Error: len(badFiles) should be 2 but is %d\n", len(badFiles))
	}

	if !reflect.DeepEqual(badFiles, []string{"test/bad1.go", "test/bad2.go"}) {
		t.Error("Error: badfiles should be [test/bad1.go, test/bad2.go] but is ", badFiles)
	}
}

func TestRetVal(t *testing.T) {
	setup()

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

func TestIgnoreStrToMap(t *testing.T) {
	for input, expected := range map[string]map[string]bool{
		"a": {
			"a": true,
		},
		"a/": {
			"a": true,
		},
		"a,b,c": {
			"a": true,
			"b": true,
			"c": true,
		},
		"a/,b/": {
			"a": true,
			"b": true,
		},
		"a,b/": {
			"a": true,
			"b": true,
		},
	} {
		if val := ignoreStrToMap(input); !reflect.DeepEqual(expected, val) {
			t.Errorf("Error: expected %s to produce %s but got %s\n",
				input,
				fmt.Sprint(expected),
				fmt.Sprint(val),
			)
		}
	}
}

func TestIgnoreWalk(t *testing.T) {
	setup()
	ignoreStr = "test/"
	ignoreMap = ignoreStrToMap(ignoreStr)

	err := filepath.Walk(rootDir, fileCheck)
	if err != nil {
		t.Errorf("Error walking directory: %s\n", err.Error())
	}

	if len(badFiles) != 0 {
		t.Errorf("Error: len(badFiles) should be 0 but is %d\n", len(badFiles))
	}

	if ret != 0 {
		t.Errorf("Error: ret should be 0 but is %d\n", ret)
	}
}
