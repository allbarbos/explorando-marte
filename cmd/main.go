package main

import (
	"explorando-marte/explore"
	"explorando-marte/utils"
	"fmt"
	"log"
)

var filePath string = "input.csv"

func main() {
	file := utils.OpenFile(filePath)
	defer file.Close()

	probes, err := explore.Init(file)

	if err != nil {
		log.Panic(err)
	}

	for _, p := range probes {
		fmt.Println(p.GetPosition())
	}
}
