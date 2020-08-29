package main

import (
	"explorando-marte/explore"
	"explorando-marte/utils"
	"fmt"
	"log"
	"path/filepath"
)

var filePath string = "input.csv"

func main() {
	file := utils.OpenFile(filepath.Clean(filePath))

	probes, err := explore.Init(file)

	if err != nil {
		log.Panic(err)
	}

	for _, p := range probes {
		fmt.Println(p.GetPosition())
	}

	_ = file.Close()
}
