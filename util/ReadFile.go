package utils

import "os"

func ReadFile(path string) string {
	buffer, err := os.ReadFile(path)
	if err != nil {
		panic("no input file")
	}

	return string(buffer)
}
