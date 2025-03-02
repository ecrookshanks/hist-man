/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ecrookshanks/hist-man/hist"
	"github.com/spf13/cobra"
)

var showBytes bool
var showDupes bool
var showUnique bool

// sizeCmd represents the size command
var sizeCmd = &cobra.Command{
	Use:   "size",
	Short: "Display the size of the history file.",
	Long: `Defaults to number of lines.
	
	Flags for size in bytes (b), unique entries (u), and dup entries (d).
	
	Examples:
	  hist size -b 		// displays the number of lines and the size in bytes.
	  hist size -bu		// displays the number of lines, size in bytes, and the number of unique lines.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		results, err := hist.GetBashFileStats()
		if err != nil {
			fmt.Println("Error getting bash history stats!")
			return
		}

		fmt.Println("Total lines: ", results.Lines)

		if showBytes {
			fmt.Println("File Size (bytes): ", results.Size)
		}
		if showDupes {
			fmt.Println("Duplicate lines: ", results.Dups)
		}

		if showUnique {
			fmt.Println("Unique Lines: ", results.Unique)
		}

	},
}

func init() {
	rootCmd.AddCommand(sizeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	sizeCmd.Flags().BoolVarP(&showBytes, "bytes", "b", false, "Report the size of the file in bytes.")
	sizeCmd.Flags().BoolVarP(&showDupes, "dupes", "d", false, "Report the number of duplicate lines.")
	sizeCmd.Flags().BoolVarP(&showUnique, "unique", "u", false, "Report the number of unique lines.")
}
