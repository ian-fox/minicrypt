package internal

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {
	var grid Grid
	grid.SetWord(0, "abcde")
	grid.SetWord(1, "fghij")
	
	assert.Equal(t, []string{"abcde", "fghij"}, grid.Horizontals(), 
		"Received unexpected output from Grid.Horizontals()")

	assert.Equal(t, []string{"af", "bg", "ch", "di", "ej"}, grid.Verticals(), 
		"Received unexpected output from Grid.Verticals()")

	assert.Equal(t, "abcde\nfghij\n     \n     \n     \n", grid.String(), 
		"Received unexpected output from Grid.String()")

	grid.ClearRow(0)
	assert.Equal(t, "     \nfghij\n     \n     \n     \n", grid.String(),
		"ClearRow did not clear a row.")
}