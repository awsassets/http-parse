package util

import "runtime"

func DeleteLastChar(data string) string {
	r := []rune(data)
	return string(r[:len(r)-1])
}

func GetLineSep() string {
	sysType := runtime.GOOS
	var lineSep string
	if sysType == "windows" {
		lineSep = "\r\n"
	} else {
		lineSep = "\n"
	}
	return lineSep
}
