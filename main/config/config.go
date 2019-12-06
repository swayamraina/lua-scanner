package config


// @author 	swayam.raina
// @dated 	05.12.2019



// This structure holds the config for the scanner
//
type Config struct {

	// this property acts as a switch for external
	// lua script discovery scan
	ExternalScan bool

	// this property decides the refresh rate for
	// the external scan of lua scripts
	RefreshRate int

	// this property is used by the scanner to
	// look for all the lua scripts in the file
	// system for discovery
	ExternalPath string

	// this property acts as a switch for internal
	// lua script discovery scan
	InternalScan bool

}