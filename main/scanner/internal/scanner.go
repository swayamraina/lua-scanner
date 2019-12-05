package internal


// @author 	swayam.raina
// @dated 	05.12.2019

import (
	"os"
	"path/filepath"
	"lua-scanner/main/scanner/utils"
)

// This function uses function currying.
// When ever we see a ".lua" file, we add
// it into the discovery result
func visit(files *[]string) filepath.WalkFunc {
	return func (path string, info os.FileInfo, err error) error {
		if LUA_EXTENSION == filepath.Ext(path) {
			*files = append(*files, path)
		}
		return nil
	}
}

// Walk function will be called for every
// file seen in the filepath "path"
func Scan (path string) []string {
	var files []string
	filepath.Walk(path, visit(&files))
	return files
}