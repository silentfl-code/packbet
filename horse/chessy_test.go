package chessy

import (
	"testing"
)

var tests_for_coordinates = []struct {
	input         string
	expectedError bool
	X, Y          int
}{
	{"a1", false, 0, 0},
	{"b1", false, 1, 0},
	{"c1", false, 2, 0},
	{"a8", false, 0, 7},
	{"d4", false, 3, 3},
	{"h1", false, 7, 0},
	{"h8", false, 7, 7},

	{"", true, 0, 0},
	{"a", true, 0, 0},
	{"a9", true, 0, 0},
	{"h98", true, 0, 0},
	{"k8", true, 0, 0},
	{"123", true, 0, 0},
}

func Test_NotationToCoords(t *testing.T) {
	for _, test := range tests_for_coordinates {
		c, err := NotationToCoords(test.input)
		if test.expectedError && err == nil || !test.expectedError && err != nil {
			t.Fatal(err)
		}
		if c.X != test.X || c.Y != test.Y {
			t.Fatal("Incorrect transformation")
		}
	}
}

func equal_moves(moves1, moves2 []string) bool {
	moves := make(map[string]bool)
	for _, move := range moves1 {
		moves[move] = true
	}

	for _, move := range moves2 {
		if _, ok := moves[move]; !ok {
			return false
		}
		delete(moves, move)
	}

	return len(moves) == 0
}

func Test_HorseMoves(t *testing.T) {
	tests := []struct {
		coord  string
		result []string
	}{
		{"a1", []string{"b3", "c2"}},
		{"h8", []string{"g6", "f7"}},
		{"b1", []string{"c3", "d2", "a3"}},
		{"b7", []string{"d8", "d6", "c5", "a5"}},
		{"c1", []string{"d3", "e2", "a2", "b3"}},
		{"c6", []string{"d8", "e7", "e5", "d4", "b4", "a5", "a7", "b8"}},
		{"f2", []string{"g4", "h3", "h1", "d1", "d3", "e4"}},
		{"g6", []string{"h8", "h4", "f4", "e5", "e7", "f8"}},
		{"h3", []string{"g1", "f2", "f4", "g5"}},
		{"h7", []string{"g5", "f6", "f8"}},
	}

	for i, test := range tests {
		coord, _ := NotationToCoords(test.coord)
		result := HorseMoves(coord)
		if !equal_moves(result, test.result) {
			t.Fatal("Wrong output for test #%d\n", i)
		}
	}
}
