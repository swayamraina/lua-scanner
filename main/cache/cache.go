
// @author 	swayam.raina
// @dated 	05.12.2019


package cache


//
//
type Cache map[string]string


//
//
func (cache Cache) get (key string) string {
	return cache[key]
}


//
//
func (cache Cache) set (key string, value string) {
	cache[key] = value
}