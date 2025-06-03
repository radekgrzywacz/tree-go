package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

const (
	Green = "\033[32m"
	Reset = "\033[0m"
)

func main() {
	var excludedExtensions []string
	var excludeHidden bool
	pflag.StringSliceVar(&excludedExtensions, "extensions", []string{}, "Files extensions to skip while going through files")
	pflag.BoolVar(&excludeHidden, "hidden", true, "Set false if you want to include hidden files")
	pflag.CommandLine.Parse(os.Args[1:])

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not find an execution path: %s", err)
	}

	var b strings.Builder
	var dirs int
	var filesAmount int

	spaces := ""
	b.WriteString(".\n")
	filePathWalkDir(wd, &b, spaces, &dirs, &filesAmount, excludedExtensions, excludeHidden)
	fmt.Print(b.String())
	fmt.Printf("\n%v directories, %v files\n", dirs, filesAmount)
}

func filePathWalkDir(root string, b *strings.Builder, prefix string, dirs, filesAmount *int, excludedExtensions []string, excludeHidden bool) error {
	files, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	files = removeUnwantedFiles(files, excludedExtensions, excludeHidden)

	for i, file := range files {
		isLast := i == len(files)-1
		var connector string
		if isLast {
			connector = "└── "
		} else {
			connector = "├── "
		}

		if file.IsDir() {
			b.WriteString(fmt.Sprintf("%s%s%s%s%s\n", prefix, connector, Green, file.Name(), Reset))
			newRoot := root + "/" + file.Name()
			newPrefix := prefix
			if isLast {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}
			filePathWalkDir(newRoot, b, newPrefix, dirs, filesAmount, excludedExtensions, excludeHidden)
			*dirs++
		}
		b.WriteString(fmt.Sprintf("%s%s%s\n", prefix, connector, file.Name()))
		*filesAmount++
	}
	return nil
}

func hasSuffixIn(word string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(word, suffix) {
			return true
		}
	}
	return false
}

func removeUnwantedFiles(files []os.DirEntry, excludedExtensions []string, excludeHidden bool) []os.DirEntry {
	filtered := files[:0]

	for _, file := range files {
		name := file.Name()

		if excludeHidden && strings.HasPrefix(name, ".") {
			continue
		}

		if hasSuffixIn(name, excludedExtensions) {
			continue
		}

		filtered = append(filtered, file)
	}

	return filtered
}
