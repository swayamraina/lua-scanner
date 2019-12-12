/*

	This service is the scan orchestrator which uses internal, external,
	database, document scanners to get all the data and build the
	required in-memory cache.

	@author 	swayam.raina
	@dated 		05.12.2019

 */


package scanner


import (
	"github.com/swayamraina/lua-scanner/main/cache"
	"github.com/swayamraina/lua-scanner/main/config"
	"github.com/swayamraina/lua-scanner/main/scanner/internal"
	"github.com/swayamraina/lua-scanner/main/utils"
)



/*
	This structure defines the main scanner service which does the
	actual heavy lifting of file-system scanning
 */
type Scanner struct {
	Config config.Config
}



/*
	This method does an,
		1. internal scan
	to discover lua scripts in the project

	return:		discovery set (all the resource file paths)
 */
func (scanner *Scanner) Scan () []string {
	var scripts []string
	config := scanner.Config
	if config.InternalScan {
		scripts = internal.Scan(scanner.Config.InternalPaths)
	}
	return scripts
}


/*
	This method accepts the scanned script paths and builds an
	in-memory cache for services requests rather than directly
	using embedded scripts

	scriptPaths:		discovery set (all the resource file paths)
	return:  			the in-memory cache instance
 */
func (scanner *Scanner) LoadInCache (scriptPaths []string) cache.Cache {
	var cache = make(cache.Cache)
	for _, scriptPath := range scriptPaths {
		filename := utils.GetFileName(scriptPath)
		content := utils.GetContents(scriptPath)
		cache.Set(filename, content)
	}
	return cache
}