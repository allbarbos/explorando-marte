package explore

import (
	"encoding/csv"
	co "explorando-marte/coordinates"
	"explorando-marte/plateau"
	p "explorando-marte/probe"
	"explorando-marte/utils"
	"log"
	"os"
	"strconv"
)

func Init(file *os.File) []p.Probe {
	r := csv.NewReader(file)
	c := 1

	var probes []p.Probe
	limit := plateau.NewLimit()

	for {
		rec := utils.ReadLine(r)
		if rec == nil {
			break
		}

		if c == 1 {
			lx, ly := readLimits(rec)
			limit.SetUpper(lx, ly)
			c++
			continue
		}

		card, route, x, y := readCoordinates(r, rec)
		probe := p.NewProbe(card, route, x, y)

		for _, r := range probe.GetRoute() {
			cmd := string(r)

			x, y, card := co.Coordinates(cmd, probe.GetCardinal())

			probe.Move(card, x, y)
			limit.Validate(probe)
		}

		probes = append(probes, probe)
		c += 2
	}

	return probes
}

func readCoordinates(r *csv.Reader, rec []string) (string, string, int, int) {
	x, err := strconv.Atoi(rec[0])
	y, err := strconv.Atoi(rec[1])
	card := rec[2]

	if err != nil {
		log.Fatal(err)
	}

	rec = utils.ReadLine(r)
	route := rec[0]

	if err != nil {
		log.Fatal(err)
	}

	return card, route, x, y
}

func readLimits(rec []string) (int, int) {
	x, err := strconv.Atoi(rec[0])
	y, err := strconv.Atoi(rec[1])

	if err != nil {
		log.Fatal(err)
	}

	return x, y
}
