
// @author 	swayam.raina
// @dated 	06.12.2019


package main


import (
	"github.com/swayamraina/lua-scanner/main/cache"
	"github.com/swayamraina/lua-scanner/main/config"
	"github.com/swayamraina/lua-scanner/main/scanner"
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