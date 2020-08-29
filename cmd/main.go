package main

import (
	"explorando-marte/explore"
	"explorando-marte/utils"
	"fmt"
)

func main() {
	file := utils.OpenFile("input.csv")
	defer file.Close()

	probes := explore.Init(file)

	for _, p := range probes {
		fmt.Println(p.GetPosition())
	}
}
