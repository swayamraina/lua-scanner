/*

	This file contains internal scanning logic i.e.
	the internal scanner will scan lua scripts which are
	embedded inside the code repository.

	@author 	swayam.raina
	@dated 		05.12.2019

 */


package internal


import (
	"github.com/swayamraina/lua-scanner/main/scanner/utils"
	"os"
	"path/filepath"
)



/*
	This function uses function currying to scan over the
	input directories and search for '.lua' files and into
	the discovery set.

	files:		the final result set containing script paths
 */
func visit(files *[]string) filepath.WalkFunc {
	return func (path string, info os.FileInfo, err error) error {
		if utils.LUA_EXTENSION == filepath.Ext(path) {
			*files = append(*files, path)
		}
		return nil
	}
}


/*
	This function starts the scan over the requested scan-paths
	and adds the 'lua' script paths to the discoverable set.
 */
func Scan (paths []string) []string {
	var files []string
	for _, path := range paths {
		filepath.Walk(path, visit(&files))
	}
	return files
}