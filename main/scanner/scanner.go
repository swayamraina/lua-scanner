
// @author 	swayam.raina
// @dated 	05.12.2019


package scanner

import (
	"github.com/swayamraina/lua-scanner/main/cache"
	"github.com/swayamraina/lua-scanner/main/config"
	"github.com/swayamraina/lua-scanner/main/scanner/internal"
	"github.com/swayamraina/lua-scanner/main/utils"
)


// This structure defines the main scanner
// service which does the actual heavy
// lifting of file-system scanning
//
type Scanner struct {
	Config config.Config
}


// This method does an internal and external scan
// to discover lua scripts in the project
//
func (scanner *Scanner) Scan () []string {
	var scripts []string
	config := scanner.Config
	if config.InternalScan {
		scripts = internal.Scan(scanner.Config.InternalPaths)
	}
	return scripts
}


//
//
func (scanner *Scanner) LoadInCache (scriptPaths []string) cache.Cache {
	var cache = make(cache.Cache)
	for _, scriptPath := range scriptPaths {
		filename := utils.GetFileName(scriptPath)
		content := utils.GetContents(scriptPath)
		cache.Set(filename, content)
	}
	return cache
}