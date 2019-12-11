/*

	This file is the defines the config for the
	'lua-scanner'.

	@author 	swayam.raina
	@dated 		05.12.2019

 */


package config



/*
	This data structure stores the config for the
	resource scanner.
	The config object is to provided by the application
	owners into the 'processor.Process()' function call
	for the scanner to do the resource discovery.
 */
type Config struct {

	/*
		This property is used by the scanner to look
		for all the lua scripts in the provided directory
		for discovery.
	 */
	InternalPaths []string


	/*
		This property acts as a switch for internal
		lua script discovery scan
	 */
	InternalScan bool

}