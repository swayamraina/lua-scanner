
// @author 	swayam.raina
// @dated 	05.12.2019


package internal


import (
	"lua-scanner/main/scanner/utils"
	"os"
	"path/filepath"
)


// This function uses function currying.
// When ever we see a ".lua" file, we add
// it into the discovery result
//
func visit(files *[]string) filepath.WalkFunc {
	return func (path string, info os.FileInfo, err error) error {
		if utils.LUA_EXTENSION == filepath.Ext(path) {
			*files = append(*files, path)
		}
		return nil
	}
}


// Walk function will be called for every
// file seen in the filepath "path"
//
func Scan (paths []string) []string {
	var files []string
	for _, path := range paths {
		filepath.Walk(path, visit(&files))
	}
	return files
}