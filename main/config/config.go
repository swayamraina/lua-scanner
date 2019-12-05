package config


// @author 	swayam.raina
// @dated 	05.12.2019



// This structure holds the config for the scanner
//
type Config struct {

	// this property acts as a switch for external
	// lua script discovery scan
	externalScan bool

	// this property decides the refresh rate for
	// the external scan of lua scripts
	refreshRate int

	// this property is used by the scanner to
	// look for all the lua scripts in the file
	// system for discovery
	externalPath string

	// this property acts as a switch for internal
	// lua script discovery scan
	internalScan bool

}