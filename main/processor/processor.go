/*

	This file declares a public method which is
	used by the user application to pitch in with the
	scan config to get the resource cache.

	@author 	swayam.raina
	@dated 		06.12.2019

 */


package processor


import (
	"github.com/swayamraina/lua-scanner/main/cache"
	"github.com/swayamraina/lua-scanner/main/config"
	"github.com/swayamraina/lua-scanner/main/scanner"
)



/*
	This function is the 'open to world' function for this
	lib.
	Application needs to pass in the scanner config which
	is used by the scanner to create a cache object which
	stores the resource contents by its name.

	config:			the scanner config
	return: 		the cache object storing the resource data
 */
func Process (config config.Config) cache.Cache {
	scanner := scanner.Scanner {
		 Config: config,
	}
	scripts := scanner.Scan()
	return scanner.LoadInCache(scripts)
}