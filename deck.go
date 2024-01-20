package main

// A Set or mini-Set deck.
type Deck interface {
	// Number of cards remaining in the deck.
	Size() int

	// Returns the card at the indicated position in the deck.
	// The index is valid over [0, Size) and
	// if index is not valid, nil is returned.
	At(index int) SetCard

	// Returns the deck containing all cards in the same order except
	// without the card at the provided index.
	// If the idnex is out of bounds, nil is returned.
	RemoveAt(index int) Deck

	// Returns true if the three cards form a Set according to the rules.
	IsValidSet(a, b, c SetCard) bool

	// Returns the card that will complete the set for the pair of cards.
	// When neither i nor j have any UNK_* properties, and are not the same,
	// there will always be only one choice possible for the third card.
	// If either card has an UNK_* property, the returned value is nil.
	CompletesSet(i, j SetCard) SetCard
}

// Draw the next card based on the deck's current ordering.
func DrawNext(deck Deck) (SetCard, Deck) {
	return deck.At(0), deck.RemoveAt(0)
}

// type definition for Intn(k)-like functions.  Should follow the semantics of
// rand.Intn where the codomain is in [0, k); i.e., `0 <= {selected} < k`.
type IntnFun func(int) int

// Package-private method for shuffling the deck (or any arbitrary slice)
// of Set cards.  Efficiently and thoroughly produces a random ordering as
// determined by rng, a ranged-integer random number generator.
func shuffleInPlace[cardType miniSetCard | fullSetCard](cards []cardType, rng IntnFun) {
	// Uses the Fisher-Yates shuffling algorithm, iteratively swapping the
	// final value with a randomly selected index into the deck and repeating
	// until the stack to be shuffled is only that final card.  This variant
	// of the original is also called "Algorithm P (Shuffling)" by Knuth.
	for size := len(cards); size > 1; size-- {
		index := rng(size)
		if index == size-1 {
			continue
		}
		cards[size-1], cards[index] = cards[index], cards[size-1]
	}
	// This could be reformulated as a non-modifying shuffle by using the
	// "inside-out" algorithm, but the way we are using decks here it is fine
	// to leave that for another time.
}

// Searches the deck for the indicated card and returns a new deck without it.
func RemoveCard(deck Deck, card SetCard) (Deck, error) {
	return nil, nil
}

// Common type for mini-Set and full Set decks.
type SetCard interface {
	Count() Count
	Color() Color
	Shape() Shape
	Shading() Shading

	Equals(other SetCard) bool
}

type Count byte
type Color byte
type Shape byte
type Shading byte

const (
	UNK_COUNT Count = 0
	ONE       Count = 1
	TWO       Count = 2
	THREE     Count = 3

	UNK_COLOR Color = 0
	RED       Color = 1
	GREEN     Color = 2
	PURPLE    Color = 3

	UNK_SHAPE Shape = 0
	ROUND     Shape = 1
	POINTY    Shape = 2
	WIGGLE    Shape = 3

	UNK_SHADING Shading = 0
	HOLLOW      Shading = 1
	PARTIAL     Shading = 2
	SOLID       Shading = 3
)

// Determine what the third value would be for a pair of a given property.
//
// If the two values are the same, return the same; if the two values are
// different, return the one that is missing.
func other[Value Count | Color | Shape | Shading](i, j Value) Value {
	if i == 0 || j == 0 {
		return (Value)(0)
	}
	if i == j {
		return i
	}
	return (Value)(6 - i - j)
}
