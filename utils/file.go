package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(filepath.Clean(path))

	if err != nil {
		log.Panic("couldn't open file", err)
	}

	return file
}

func ReadLine(r *csv.Reader) []string {
	l, err := r.Read()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		log.Panic(err)
	}
	return l
}
