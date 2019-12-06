
// @author 	swayam.raina
// @dated 	05.12.2019


package scanner


import (
	"lua-scanner/main/cache"
	"lua-scanner/main/config"
	"lua-scanner/main/scanner/internal"
)


// This structure defines the main scanner
// service which does the actual heavy
// lifting of file-system scanning
//
type Scanner struct {
	config config.Config
}


// This method does an internal and external scan
// to discover lua scripts in the project
//
func (scanner *Scanner) Scan () []string {
	var scripts []string
	config := scanner.config
	if config.InternalScan {
		scripts = internal.Scan(scanner.config.InternalPaths)
	}
	return scripts
}


//
//
func (scanner *Scanner) LoadInCache (scriptPaths []string) cache.Cache {
	var cache = make(cache.Cache)

	return cache
}