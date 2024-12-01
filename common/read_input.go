package common

import (
	"io"
	"os"
)

func ReadInput(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	intput, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(intput)
}
