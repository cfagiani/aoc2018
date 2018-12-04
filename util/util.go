// Package util provides utility functions that are not specific to any single problem from AoC.
package util

import (
	"io/ioutil"
	"log"
)

//Checks the error argument and, if it is not nil, it will log the msg passed in. If isFatal is true, the log will be
//written as Fatal which will cause exit(1) to be called.
func CheckError(err error, msg string, isFatal bool) bool {
	if err != nil {
		if isFatal {
			log.Fatal(msg, err)
		} else {
			log.Println(msg)
		}
		return true
	}
	return false
}

// Returns the full contents of a file as a string. If the file cannot be read, it will log a Fatal error and exit the program.
func ReadFileAsString(fname string) string {
	dat, err := ioutil.ReadFile(fname)
	CheckError(err, "Could not read file", true)
	return string(dat)
}

