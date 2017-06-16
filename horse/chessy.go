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
		Coordinate{1, 2},
		Coordinate{2, 1},
		Coordinate{2, -1},
		Coordinate{1, -2},
		Coordinate{-1, -2},
		Coordinate{-1, 2},
		Coordinate{-2, -1},
		Coordinate{-2, 1},
	}
)

func NotationToCoords(coords string) (c Coordinate, err error) {
	if !isCoords(coords) {
		err = errors.New("Wrong coordinate")
		return
	}

	return Coordinate{int(coords[0] - 'a'), int(coords[1] - '1')}, nil
}

func CoordsToNotation(c Coordinate) string {
	return fmt.Sprintf("%c%d", c.X+97, c.Y+1)
}

func HorseMoves(coords Coordinate) (res []string) {
	for _, delta := range horseSteps {
		new_coords := Coordinate{coords.X + delta.X, coords.Y + delta.Y}
		if new_coords.X >= 0 && new_coords.X <= 7 &&
			new_coords.Y >= 0 && new_coords.Y <= 7 {
			res = append(res, CoordsToNotation(new_coords))
		}
	}

	return
}
