package explore

import (
	"encoding/csv"
	co "explorando-marte/coordinates"
	"explorando-marte/plateau"
	p "explorando-marte/probe"
	"explorando-marte/utils"
	"os"
	"strconv"
)

func Init(file *os.File) ([]p.Probe, error) {
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
			lx, ly, err := readLimits(rec)
			if err != nil {
				return nil, err
			}

			limit.SetUpper(lx, ly)
			c++
			continue
		}

		card, route, x, y, err := readCoordinates(r, rec)

		if err != nil {
			return nil, err
		}

		probe := p.NewProbe(card, route, x, y)

		for _, r := range probe.GetRoute() {
			cmd := string(r)

			x, y, card := co.Coordinates(cmd, probe.GetCardinal())

			probe.Move(card, x, y)

			err := limit.Validate(probe)

			if err != nil {
				return nil, err
			}
		}

		probes = append(probes, probe)
		c += 2
	}

	return probes, nil
}

func readCoordinates(r *csv.Reader, rec []string) (string, string, int, int, error) {
	x, err := strconv.Atoi(rec[0])
	if err != nil {
		return "", "", 0, 0, err
	}

	y, err := strconv.Atoi(rec[1])
	if err != nil {
		return "", "", 0, 0, err
	}

	card := rec[2]

	rec = utils.ReadLine(r)
	route := rec[0]

	return card, route, x, y, nil
}

func readLimits(rec []string) (int, int, error) {
	x, err := strconv.Atoi(rec[0])

	if err != nil {
		return 0, 0, err
	}

	y, err := strconv.Atoi(rec[1])

	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}
