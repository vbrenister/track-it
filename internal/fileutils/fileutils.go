package fileutils

import (
	"os"
)

func CreateDirIfNotExists(dir string) {
	if !IsExist(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}
	}
}

func CreateFileIfNotExists(fileName string) {
	if !IsExist(fileName) {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}
}

func IsExist(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

func MustOpen(fileName string) *os.File {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	return f
}
