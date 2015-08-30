package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	rootDir  string
	badFiles []string
	ret      int
)

func main() {
	flag.StringVar(&rootDir, "d", ".", "Directory to search")
	flag.Parse()

	err := filepath.Walk(rootDir, fileCheck)
	if err != nil {
		fmt.Printf("Error checking: %s\n", err.Error())
		fmt.Printf("exit status %d\n", ret)
		os.Exit(ret)
	}

	if len(badFiles) != 0 {
		fmt.Println("Improperly formatted files detected:")

		for _, name := range badFiles {
			fmt.Printf("- %s\n", name)
		}

		fmt.Printf("\nexit status %d\n", ret)
		os.Exit(ret)
	}
}

// Go through all go source files rooted at -d, if find any improperly
// formatted files append them to the list
func fileCheck(path string, info os.FileInfo, err error) error {
	if err != nil {
		ret = 1
		return err
	}

	if info.IsDir() {
		return nil
	}

	if filepath.Ext(path) == ".go" {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		formatted, err := format.Source(content)
		if err != nil {
			return err
		}

		if !bytes.Equal(content, formatted) {
			ret = 2
			badFiles = append(badFiles, path)
		}
	}

	return nil
}
