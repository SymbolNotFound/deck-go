package main

import (
	"fmt"
	"strings"
)

const COUNT_MINISET int = 3 * 3 * 3

// Creates an unshuffled deck of mini-Set cards.
func NewMiniSetDeck() Deck {
	deck := miniSetDeck{make([]miniSetCard, COUNT_MINISET)}
	i := 0
	for _, count := range []Count{ONE, TWO, THREE} {
		for _, color := range []Color{RED, GREEN, PURPLE} {
			for _, shape := range []Shape{ROUND, POINTY, WIGGLE} {
				deck.Cards[i] = miniSetCard{count, color, shape}
				i++
			}
		}
	}
	return deck
}

//
// Single card definition for mini-Set
//

// Struct representation of a single card from the mini-Set deck.
type miniSetCard struct {
	count Count
	color Color
	shape Shape
}

// The number of symbols on the card.  Readonly property outside of the package.
func (card *miniSetCard) Count() Count { return card.count }

// The symbol's color on the card.  Readonly property outside of the package.
func (card *miniSetCard) Color() Color { return card.color }

// The symbol's color on the card.  Readonly property outside of the package.
func (card *miniSetCard) Shape() Shape { return card.shape }

// In mini-Set there are only three properties, shading is indeterminate
// (frontends may choose to use whatever shading they want, as long as it is consistent)
func (card *miniSetCard) Shading() Shading { return UNK_SHADING }

// Returns true if the two card pointers represent the same card.
// (i.e., have the same values for each property).
func (card *miniSetCard) Equals(other SetCard) bool {
	if other, ok := other.(*miniSetCard); ok {
		return (card.count == other.count && card.color == other.color && card.shape == other.shape)
	}
	return false
}

func (card *miniSetCard) String() string {
	return strings.Join([]string{
		fmt.Sprint(card.count),
		fmt.Sprint(card.color),
		fmt.Sprint(card.shape),
	}, " ")
}

//
// Deck of mini-Set cards
//

type miniSetDeck struct {
	Cards []miniSetCard
}

func (deck miniSetDeck) Size() int {
	return len(deck.Cards)
}

func (deck miniSetDeck) At(index int) SetCard {
	if index < 0 || index > len(deck.Cards) {
		return nil
	}
	return &(deck.Cards[index])
}

func (deck miniSetDeck) RemoveAt(index int) Deck {
	if index < 0 || index >= len(deck.Cards) {
		return nil
	}
	if index == 0 {
		return miniSetDeck{deck.Cards[1:]}
	}
	if index == len(deck.Cards)-1 {
		return miniSetDeck{deck.Cards[:index]}
	}
	return miniSetDeck{Cards: append(deck.Cards[:index], deck.Cards[index+1:]...)}
}

// Returns true if the three cards together form a Set.
//
// This is true if for each property (count, color, shape) either all three
// cards have the same value or all three are different values.
func (deck miniSetDeck) IsValidSet(i, j, k SetCard) bool {
	return k.(*miniSetCard).Equals(deck.CompletesSet(i, j))
}

// Returns the third card that completes the set for the indicated cards.
//
// For a pair of (non-equal, completely defined) cards, there is only one card.
// Returns nil if the pair of cards are equal or not completely defined.
func (deck miniSetDeck) CompletesSet(i, j SetCard) SetCard {
	if i.(*miniSetCard).Equals(j) {
		return nil
	}
	card := miniSetCard{}

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

	return &card
}
