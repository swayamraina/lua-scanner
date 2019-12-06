
// @author 	swayam.raina
// @dated 	05.12.2019


package config


// This structure holds the config for the scanner
//
type Config struct {

	// this property is used by the scanner to
	// look for all the lua scripts in the
	// project for discovery
	//
	InternalPaths []string

	// this property acts as a switch for internal
	// lua script discovery scan
	//
	InternalScan bool

}