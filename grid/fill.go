package grid

import (
	"errors"
	"math/rand"
	"time"

	"github.com/ian-fox/minicrypt/wordutil"
)

// randomArray returns a shuffled list of indices.
func randomArray(length int) []int {
	arr := make([]int, length)

	for i := 0; i < length; i++ {
		arr[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(length, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	return arr
}

// Return true if a branch is invalid and should be pruned.
func prune(grid *Grid, prefixes wordutil.PrefixTree) bool {
	for _, word := range grid.Verticals() {
		if !prefixes.Check(word) {
			return true
		}
	}

	return false
}

// fillOne is a helper for fill. It places one word in the grid.
func fillOne(grid *Grid, prefixes wordutil.PrefixTree, words []string, level int) bool {
	if level == 5 {
		return true
	}

	order := randomArray(len(words))

	// Try all words in a random order
	for _, i := range order {
		grid.SetWord(level, words[i])

		// is this branch worth pursuing?
		if prune(grid, prefixes) {
			continue
		}

		if fillOne(grid, prefixes, words, level+1) {
			return true
		}
	}

	// No word from this point can create a valid grid.
	grid.ClearRow(level)
	return false
}

// Fill will create a valid 5x5 grid.
func Fill(words []string) (*Grid, error) {
	// Make the tree.
	prefixes := wordutil.NewTreeNode()
	for _, word := range words {
		prefixes.AddWord(word)
	}

	grid := Grid{}

	// Try to fill the grid
	if fillOne(&grid, prefixes, words, 0) {
		return &grid, nil
	}

	return nil, errors.New("could not construct valid grid with wordlist")
}
