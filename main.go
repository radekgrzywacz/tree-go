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

func filePathWalkDir(root string, b *strings.Builder, spaces string) error {
	files, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	files = removeHiddenFiles(files)

	if len(files) == 0 {
		return nil
	}

	for idx, file := range files {
		var formattedFile string
		if idx < len(files)-1 {
			formattedFile = fmt.Sprintf("%s├── %s%s%s\n", spaces, Green, file.Name(), Reset)
		} else {
			formattedFile = fmt.Sprintf("%s└── %s%s%s\n", spaces, Green, file.Name(), Reset)
		}

		b.WriteString(formattedFile)
		if file.IsDir() {
			spaces += strings.Repeat(" ", 4)
			if len(spaces) > 0 {
				spaces = replaceAtIndex(spaces, '│', 0)
			}
			newDir := fmt.Sprintf("%s/%s", root, file.Name())
			filePathWalkDir(newDir, b, spaces)
			if len(spaces) > 6 {
				spaces = spaces[:len(spaces)-4]
			} else {
				spaces = ""
			}
		}
	}
	if len(spaces) > 4 {
		spaces = spaces[:len(spaces)-4]
	} else {
		spaces = ""
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

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
