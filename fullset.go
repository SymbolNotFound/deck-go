package main

import (
	"fmt"
	"strings"
)

const COUNT_FULLSET int = 3 * 3 * 3 * 3

// Creates an unshuffled deck of mini-Set cards.
func NewFullSetDeck() Deck {
	deck := fullSetDeck{make([]fullSetCard, COUNT_FULLSET)}
	i := 0
	for _, count := range []Count{ONE, TWO, THREE} {
		for _, color := range []Color{RED, GREEN, PURPLE} {
			for _, shape := range []Shape{ROUND, POINTY, WIGGLE} {
				for _, shading := range []Shading{HOLLOW, PARTIAL, SOLID} {
					deck.Cards[i] = fullSetCard{count, color, shape, shading}
					i++
				}
			}
		}
	}
	return deck
}

//
// Single card definition for the full Set deck
//

// Struct representation of a single card from the mini-Set deck.
type fullSetCard struct {
	count   Count
	color   Color
	shape   Shape
	shading Shading
}

// The number of symbols on the card.  Readonly property outside of the package.
func (card *fullSetCard) Count() Count { return card.count }

// The symbol's color on the card.  Readonly property outside of the package.
func (card *fullSetCard) Color() Color { return card.color }

// The symbol's color on the card.  Readonly property outside of the package.
func (card *fullSetCard) Shape() Shape { return card.shape }

// The amount of shading in the symbols.  Readonly property outside of package.
func (card *fullSetCard) Shading() Shading { return card.shading }

// Returns true if the two card pointers represent the same card.
// (i.e., have the same values for each property).
func (card *fullSetCard) Equals(other SetCard) bool {
	if other, ok := other.(*fullSetCard); ok {
		return (card.count == other.count && card.color == other.color && card.shape == other.shape && card.shading == other.shading)
	}
	return false
}

func (card *fullSetCard) String() string {
	return strings.Join([]string{
		fmt.Sprint(card.count),
		fmt.Sprint(card.color),
		fmt.Sprint(card.shape),
		fmt.Sprint(card.shading),
	}, " ")
}

//
// Deck of mini-Set cards
//

type fullSetDeck struct {
	Cards []fullSetCard
}

func (deck fullSetDeck) Size() int {
	return len(deck.Cards)
}

func (deck fullSetDeck) At(index int) SetCard {
	if index < 0 || index > len(deck.Cards) {
		return nil
	}
	return &(deck.Cards[index])
}

func (deck fullSetDeck) RemoveAt(index int) Deck {
	if index < 0 || index >= len(deck.Cards) {
		return nil
	}
	if index == 0 {
		return fullSetDeck{deck.Cards[1:]}
	}
	if index == len(deck.Cards)-1 {
		return fullSetDeck{deck.Cards[:index]}
	}
	return fullSetDeck{Cards: append(deck.Cards[:index], deck.Cards[index+1:]...)}
}

// Returns true if the three cards together form a Set.
//
// This is true if for each property (count, color, shape) either all three
// cards have the same value or all three are different values.
func (deck fullSetDeck) IsValidSet(i, j, k SetCard) bool {
	return k.(*fullSetCard).Equals(deck.CompletesSet(i, j))
}

// Returns the third card that completes the set for the indicated cards.
//
// For a pair of (non-equal, completely defined) cards, there is only one card.
// Returns nil if the pair of cards are equal or not completely defined.
func (deck fullSetDeck) CompletesSet(i, j SetCard) SetCard {
	if i.(*fullSetCard).Equals(j) {
		return nil
	}
	card := fullSetCard{}

	card.count = other(i.Count(), j.Count())
	if card.count == UNK_COUNT {
		return nil
	}
	card.color = other(i.Color(), j.Color())
	if card.color == UNK_COLOR {
		return nil
	}
	card.shape = other(i.Shape(), j.Shape())
	if card.shape == UNK_SHAPE {
		return nil
	}
	card.shading = other(i.Shading(), j.Shading())
	if card.shading == UNK_SHADING {
		return nil
	}

	return &card
}
