package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// These should form a grid.
var gridTestData = []string{
	"chess",
	"chest",
	"hence",
	"heath",
	"eaton",
	"enter",
	"steps",
	"scope",
	"three",
	"sense",
}

// Impossible to form a grid with only these words.
var extraWords = []string{
	"abuzz",
	"affix",
	"amaze",
	"aback",
	"annex",
	"apply",
	"abbey",
	"awful",
	"above",
	"amble",
}

var results = []string{
	"chest\nheath\nenter\nscope\nsense\n",
	"chess\nhence\neaton\nsteps\nthree\n",
}

func TestFill(t *testing.T) {
	grid, err := Fill(append(gridTestData, extraWords...))
	if err != nil {
		t.Errorf("Could not create grid from test data: %s", err)
	}

	assert.Contains(t, results, grid.String())

	_, err = Fill(extraWords)
	if err == nil {
		t.Error("Expected error when making grid from bad wordlist.")
	}
}
