
// @author 	swayam.raina
// @dated 	05.12.2019


package cache


// This global variable stores all the
// scanned data and is responsible to
// serve cache hits
//
var LuaScannerCache Cache

//
//
type Cache map[string]string


//
//
func (cache *Cache) Get (key string) string {
	return (*cache)[key]
}


//
//
func (cache *Cache) Set (key string, value string) {
	(*cache)[key] = value
}