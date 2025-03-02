/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/ecrookshanks/hist-man/hist"
	"strings"

	"github.com/spf13/cobra"
)

var showAll bool

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search the history list for a specific string",
	Long: `Search the history file for a specifc string.
	If found, the entire line will be returned.
	By default, ONLY the first instance will be returned. the "--all" 
	or "-a" argument will return all instances.

	for example:

	hist-man search 'dnf'    // returns the first instance of the command with dnf if it.
	hist-man search -a 'dnf' // returns all the instances of the dnf command.
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		toSearch := args[0]
		fmt.Println("search called, looking for " + toSearch)
		results, err := hist.GetBashFileStats()
		if err != nil {
			fmt.Println("Error getting bast history stats!")
			return
		}

		for _, lineVal := range results.All {
			if strings.Contains(lineVal, toSearch) {
				fmt.Println("Found match!")
				fmt.Println(lineVal)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
