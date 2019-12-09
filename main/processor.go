
// @author 	swayam.raina
// @dated 	06.12.2019


package main


import (
	"lua-scanner/main/cache"
	"lua-scanner/main/config"
	"lua-scanner/main/scanner"
)


//
//
func process (config config.Config) cache.Cache {
	scanner := scanner.Scanner {
		 Config: config,
	}
	scripts := scanner.Scan()
	return scanner.LoadInCache(scripts)
}