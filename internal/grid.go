package internal

import "strings"

// Grid holds a mini crossword.
type Grid [5][5]byte

// Empty is a constant for yet-to-be-filled squares.
var Empty = byte(0)

// Horizontals returns the horizontal words in a grid.
func (grid *Grid) Horizontals() []string {
	var horizontals []string

	for i := 0; i < 5; i++ {
		if (grid[i][0] == Empty) {
			break
		}
		horizontals = append(horizontals, string(grid[i][:]))
	}

	return horizontals
}

// Verticals returns the vertical words in a grid.
func (grid *Grid) Verticals() []string {
	var verticals []string

	for col := 0; col < 5; col++ {
		var word []byte
		for row := 0; row < 5; row++ {
			if grid[row][col] == Empty {
				break
			}
			word = append(word, grid[row][col])
		}
		verticals = append(verticals, string(word))
	}

	return verticals
}

// String returns a string representation of the grid.
func (grid *Grid) String() string {
	var output strings.Builder;
	output.Grow(6 * 5) // 6 characters times 5 lines
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if grid[i][j] == Empty {
				output.WriteString(" ")
			} else {
				output.WriteByte(grid[i][j])
			}
		}
		output.WriteString("\n")
	}

	return output.String()
}

// SetWord sets a word in a grid.
func (grid *Grid) SetWord(row int, word string) {
	for i, char := range word {
		grid[row][i] = byte(char)
	}
}

// ClearRow zeroes out a row of a grid.
func (grid *Grid) ClearRow(row int) {
	for i := 0; i < 5; i++ {
		grid[row][i] = byte(0)
	}
}
