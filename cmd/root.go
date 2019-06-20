package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ian-fox/minicrypt/grid"
)

var rootCmd = &cobra.Command{
	Use:  "minicrypt wordlist_file",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		defer f.Close()

		var words []string

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			word := scanner.Text()
			if word[0] == '#' || len(word) != 5 {
				continue
			}
			words = append(words, word)
		}

		if err := scanner.Err(); err != nil {
			return err
		}

		crossword, err := grid.Fill(words)
		if err != nil {
			return err
		}

		fmt.Println(crossword.String())
		return nil
	},
}

// Execute runs the application.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
