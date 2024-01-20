package main

import (
	"fmt"
	"math/rand"
)

const (
	NUM_TRIALS int = 5000000
	DRAW_LIMIT int = 19
)

// Within the rules of a reduced variant of Set, dealing cards individually
// instead of 12 at once face up, determine the odds of finding a complete Set
// in the first $k$ cards dealt, by direct selection and marginalization.
// Also calculate the likelihood of finding more than one set in a deal of $k$
// cards
func main() {
	table := make(map[int]map[int]int)

	for i := 0; i < NUM_TRIALS; i++ {
		deck := NewMiniSetDeck()
		//deck := NewFullSetDeck()
		shuffleInPlace(deck.(fullSetDeck).Cards, rand.Intn)
		for dealt := 3; dealt <= DRAW_LIMIT; dealt++ {
			count := countSets(deck, dealt)
			if _, ok := table[dealt]; !ok {
				table[dealt] = make(map[int]int)
			}
			table[dealt][count] += 1
		}
	}
	printTable(table)
}

func printTable(table map[int]map[int]int) {
	for i := 3; i <= DRAW_LIMIT; i++ {
		fmt.Printf("with %d cards drawn,\n", i)
		max_sets := 0
		for k := range table[i] {
			max_sets = max(max_sets, k)
		}

		positives := make([]int, max_sets)
		for j := 1; j <= max_sets; j++ {
			fmt.Printf("%9d ", j)
			count := table[i][j]
			for k := 0; k < j; k++ {
				positives[k] += count
			}
		}
		fmt.Println()

		var trials float64 = float64(NUM_TRIALS)
		for _, count := range positives {
			fmt.Printf(" %8.5f ", (float64(count)/trials)*100.)
		}
		fmt.Println()
		fmt.Println()
	}
}

// Returns the number of unique triples that complete a Set, will be >= 0.
func countSets(deck Deck, dealt int) int {
	count := 0
	for i := 0; i < dealt-2; i++ {
		for j := i + 1; j < dealt-1; j++ {
			for k := j + 1; k < dealt; k++ {
				if deck.IsValidSet(deck.At(i), deck.At(j), deck.At(k)) {
					count += 1
				}
			}
		}
	}
	return count
}
