package utils

import (
	"strings"
)

var ignoreDirs []string
var ignoreFiles []string
var tils map[string][]string

func init() {
	ignoreDirs = []string{".git", ".idea"}
	ignoreFiles = []string{".gitignore", "README.md"}
	tils = make(map[string][]string)
}

func AddTilFiles(path string) {
	if !isSkip(path) {
		addFiles(path)
	}
}

func GetTilFiles() map[string][]string {
	return tils
}

func addFiles(path string) {
	if strings.Contains(path,"/") {
		dir := strings.Split(path, "/")[0]
		files, ok := tils[dir]
		if ok {
			files = append(files, path)
			tils[dir] = files
		} else {
			tils[dir] = make([]string, 5)
			files = append(files, path)
			tils[dir] = files
		}
	}
}

func isSkip(path string) bool{
	for _, dir := range ignoreDirs {
		if strings.Contains(path, dir) {
			return true
		}
	}

	for _, file := range ignoreFiles {
		if strings.Contains(path, file) {
			return true
		}
	}
	return false
}
