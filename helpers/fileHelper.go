package helpers

import "os"

//FileExists function checks if a file exists in the given path and returns the bool result
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
