package card

import "sort"

type CardSuit byte

const (
	Clubs CardSuit = iota
	Diamonds
	Hearts
	Spades
)

type CardColor byte

const (
	red CardColor = iota
	black
)

type CardValue byte

const (
	_2 CardValue = iota
	_3
	_4
	_5
	_6
	_7
	_8
	_9
	_10
	J
	Q
	K
	A
	// NO Joker
)

var CardWeighBySuit = map[CardValue]byte{
	_2:  0,
	_3:  1,
	_4:  2,
	_5:  3,
	_6:  4,
	_7:  5,
	_8:  6,
	_9:  7,
	_10: 8,
	J:   9,
	Q:   10,
	K:   11,
	A:   12,
}

type Card struct {
	Value CardValue
	Color CardColor
	Suit  CardSuit
}

type Hand []Card

type By func(c1, c2 *Card) bool

func (by By) Sort(hand Hand) {
	ps := &handSorter{
		hand: hand,
		by:   by,
	}
	sort.Sort(ps)
}

type handSorter struct {
	hand Hand
	by   func(c1, c2 *Card) bool
}

func (h *handSorter) Len() int {
	return len(h.hand)
}

func (h *handSorter) Swap(i, j int) {
	h.hand[i], h.hand[j] = h.hand[j], h.hand[i]
}

func (h *handSorter) Less(i, j int) bool {
	return h.by(&h.hand[i], &h.hand[j])
}

func (hand Hand) SortBySuit() {
	suit := func(c1, c2 *Card) bool {
		return CardWeighBySuit[c1.Value] < CardWeighBySuit[c2.Value] && c1.Color < c2.Color
	}
	By(suit).Sort(hand)
}
