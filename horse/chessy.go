package chessy

import (
	"errors"
	"fmt"
	"regexp"
)

type Coordinate struct {
	X, Y int
}

var (
	isCoords = regexp.MustCompile("^[a-h][1-8]$").MatchString

	horseSteps = []Coordinate{
		Coordinate{1, -2},
		Coordinate{2, -1},
		Coordinate{2, 1},
		Coordinate{1, 2},
		Coordinate{-1, 2},
		Coordinate{-2, 1},
		Coordinate{-2, -1},
		Coordinate{-1, -2},
	}
)

func NotToCoords(coords string) (c Coordinate, err error) {
	if !isCoords(coords) {
		err = errors.New("Wrong coordinate")
		return
	}

	return Coordinate{int(coords[0] - 'a'), int(coords[1] - '1')}, nil
}

func HorseMoves(coords Coordinate) (res []string) {
	for _, delta := range horseSteps {
		new_x, new_y := coords.X+delta.X, coords.Y+delta.Y
		if new_x >= 0 && new_x <= 7 &&
			new_y >= 0 && new_y <= 7 {
			res = append(res, fmt.Sprintf("%c%d ", new_x+97, new_y))
		}
	}

	return
}
