package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	rootDir   string
	ignoreStr string
	ignoreMap map[string]bool
	badFiles  []string
	ret       int
)

// Perform checks according to command line flags
// Exit with status 1 if there was some some of error in walking / checking
// Exit with status 2 if incorrectly formatted files were detected
func main() {
	flag.StringVar(&rootDir, "d", ".", "Directory to search")
	flag.StringVar(&ignoreStr, "i", "", "Comma-separated directories to ignore (useful for vendored deps)")
	flag.Parse()

	ignoreMap = ignoreStrToMap(ignoreStr)

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
		if _, ok := ignoreMap[path]; ok {
			fmt.Printf("Skipping directory: %s\n", path)

			// Will direct filepath.Walk to not recurse down into directory
			return filepath.SkipDir
		}

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

// Take in a comma-separated string of directories to skip, return a map
// map from names to bool for easy lookup
func ignoreStrToMap(s string) map[string]bool {
	m := map[string]bool{}
	for _, dir := range strings.Split(s, ",") {
		if dir != "" {
			// Clean up trailing slashes if present
			m[filepath.Clean(dir)] = true
		}
	}
	return m
}
