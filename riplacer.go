package main

import (
	// "bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func interpretEscapes(s string) string {
	return strings.ReplaceAll(s, "\\n", "\n")
}

func replaceInFile(filePath string, findText string, replaceText string) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	updatedContent := strings.ReplaceAll(string(content), findText, replaceText)

	err = ioutil.WriteFile(filePath, []byte(updatedContent), 0644)
	return err
}

func processFiles(pattern string, findText string, replaceText string, includeHidden bool) error {
	return filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories, but process hidden files if needed
		if info.IsDir() {
			if isHiddenFile(filepath.Base(path)) && !includeHidden && path != "." {
				return filepath.SkipDir
			}
			return nil
		}

		// Check if the path matches the pattern
		relativePath, err := filepath.Rel(".", path)
		if err != nil {
			return err
		}

		match, err := filepath.Match(pattern, relativePath)
		if err != nil {
			return err
		}

		if match {
			err = replaceInFile(path, findText, replaceText)
			if err != nil {
				log.Printf("Failed to replace in file: %s, error: %s", path, err)
			}
		}

		return nil
	})
}

func isHiddenFile(fileName string) bool {
	return strings.HasPrefix(fileName, ".")
}

func main() {
	var pattern, findText, replaceText string
	var includeHidden bool

	flag.StringVar(&pattern, "p", "", "File pattern to modify")
	flag.StringVar(&findText, "f", "", "Text to find")
	flag.StringVar(&replaceText, "r", "", "Text for replacement")
	flag.BoolVar(&includeHidden, "hidden", false, "Include hidden files")

	flag.Parse()

	interpretedFindText := interpretEscapes(findText)

	if pattern == "" || interpretedFindText == "" || replaceText == "" {
		fmt.Println("Usage: replace -p <pattern> [-hidden] [-i] -f <find> -r <replace>")
		os.Exit(1)
	}

	err := processFiles(pattern, interpretedFindText, replaceText, includeHidden)
	if err != nil {
		log.Fatalf("Error processing files: %s", err)
	}
}
