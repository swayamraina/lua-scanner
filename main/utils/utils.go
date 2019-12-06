
// @author 	swayam.raina
// @dated 	06.12.2019


package utils


import (
	"io/ioutil"
	"log"
	"os"
)


//
//
func getContentsInBytes (scriptPath string) []byte {
	file, err := os.Open(scriptPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	return bytes
}


//
//
func GetContents (scriptPath string) string {
	bytes := getContentsInBytes(scriptPath)
	return string(bytes)
}


//
//
func GetFileName (scriptPath string) string {
	return "hello"
}