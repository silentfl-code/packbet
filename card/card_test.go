package card

import (
	"reflect"
	"testing"
)

func Test_SortBySuit(t *testing.T) {
	tests := []struct {
		input,
		output Hand
	}{
		{
			input: Hand{
				{_4, red, Diamonds},
				{_3, red, Diamonds},
				{_5, red, Diamonds},
				{_2, red, Diamonds},
			},
			output: Hand{
				{_2, red, Diamonds},
				{_3, red, Diamonds},
				{_4, red, Diamonds},
				{_5, red, Diamonds},
			},
		},
		{
			input: Hand{
				{_2, red, Diamonds},
				{_2, black, Diamonds},
				{_2, red, Clubs},
				{_2, black, Clubs},
			},
			output: Hand{
				{_2, red, Clubs},
				{_2, red, Diamonds},
				{_2, black, Clubs},
				{_2, black, Diamonds},
			},
		},
		{
			input: Hand{
				{_2, red, Hearts},
				{_2, red, Diamonds},
				{_2, red, Clubs},
				{_2, red, Spades},
			},
			output: Hand{
				{_2, red, Clubs},
				{_2, red, Diamonds},
				{_2, red, Hearts},
				{_2, red, Spades},
			},
		},
	}
	for _, test := range tests {
		test.input.SortBySuit()
		if reflect.DeepEqual(test.input, test.output) {
			t.Fatal("Order is wrong")
		}
	}
}
