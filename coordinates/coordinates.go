package coordinates

var side = map[string]map[string]string{
	"R": {
		"N": "E",
		"E": "S",
		"W": "N",
		"S": "W",
	},
	"L": {
		"N": "W",
		"E": "N",
		"W": "S",
		"S": "E",
	},
}

var front = map[string]map[string][]int{
	"M": {
		"N": []int{0, 1},
		"S": []int{0, -1},
		"E": []int{1, 0},
		"W": []int{-1, 0},
	},
}

// Coordinates get
func Coordinates(cmd, cardinal string) (int, int, string) {
	var x, y int
	card := cardinal

	if cmd == "M" {
		x = front[cmd][cardinal][0]
		y = front[cmd][cardinal][1]

		return x, y, card
	}

	return x, y, side[cmd][cardinal]
}
