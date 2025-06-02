package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Green = "\033[32m"
	Reset = "\033[0m"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not find an execution path: %s", err)
	}

	var b strings.Builder
	spaces := ""
	b.WriteString(".\n")
	filePathWalkDir(wd, &b, spaces)
	fmt.Print(b.String())

}

func filePathWalkDir(root string, b *strings.Builder, prefix string) error {
	files, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	files = removeHiddenFiles(files)

	for i, file := range files {
		isLast := i == len(files)-1
		var connector string
		if isLast {
			connector = "└── "
		} else {
			connector = "├── "
		}
		b.WriteString(fmt.Sprintf("%s%s%s%s%s\n", prefix, connector, Green, file.Name(), Reset))

		if file.IsDir() {
			newRoot := root + "/" + file.Name()
			newPrefix := prefix
			if isLast {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}
			filePathWalkDir(newRoot, b, newPrefix)
		}
	}
	return nil
}

func removeHiddenFiles(files []os.DirEntry) []os.DirEntry {
	filtered := files[:0]

	for _, file := range files {
		if !strings.HasPrefix(file.Name(), ".") {
			filtered = append(filtered, file)
		}
	}

	return filtered
}
