/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ecrookshanks/hist-man/hist"
	"github.com/spf13/cobra"
)

type ShowOptions struct {
	showBeginning  bool
	showLatest     bool
	showCount      int
	showDup        bool
	showUniqueVals bool
	showAll        bool
	showMaxDup     bool
}

func NewShowOptions() *ShowOptions {
	return &ShowOptions{
		showAll: true,
	}
}

var longDesc string = `Show a subset of entries in the bash (or zsh) history file. For example:

		hist show 			// show all the entries in the file.
		hist show -b		// show the beginning (oldest) 10 entries.
		hist show -n		// show the last (most recent) 10 entries.
		hist show -n -c 20	// show the last 20 entries in the file.
		hist show -u 		// show the unique entries.
		hist show -d		// show only the duplicated entries.
		hist show -dm		// show the single most duplicated entry.

		** potential additions **
		hist show -dn		// show the 10 latest duplicated entries.
		hist show -un		// show the 10 latest unique entries.
		
`

// showCmd represents the show command
func CreateShowCmd() *cobra.Command {

	o := NewShowOptions()

	var showCmd = &cobra.Command{

		Use:   "show",
		Short: "Show a selected set of entries in the history file.",
		Long:  longDesc,
		Run: func(cmd *cobra.Command, args []string) {

			results, err := hist.GetBashFileStats()
			if err != nil {
				fmt.Println("Error getting history content!")
				return
			}
			fmt.Println("**** HISTORY RESULTS ****")

			if o.showBeginning {
				fmt.Printf("*** Oldest %d enties:\n", o.showCount)
				for _, ln := range results.All[0:o.showCount] {
					fmt.Println(ln)
				}
				return
			}

			if o.showLatest {
				fmt.Printf("*** Most recent %d entries:\n", o.showCount)
				elements := len(results.All)
				for _, ln := range results.All[elements-o.showCount : elements] {
					fmt.Println(ln)
				}
				o.showAll = false
				return
			}

			if o.showDup {
				if o.showMaxDup {
					maxKey, max := hist.FindMaxDupValueAndName(results.DupCounts)

					fmt.Printf("\"%s\" is the most duplicated entry, total %d repeats", maxKey, max)
					o.showAll = false
					return
				}
				fmt.Println("*** All the duplicated entries: ")
				for _, ln := range results.DupVals {
					fmt.Println(ln)
				}
				o.showAll = false
			}

			if o.showUniqueVals {
				fmt.Println("*** All the unique entries:")
				for _, ln := range results.UniqueVals {
					fmt.Println(ln)
				}
				o.showAll = false
			}

			//
			//TODO: Is there a better way to have a "default" action? Managing a bool
			// 		seems like a bad idea.
			//
			if o.showAll {
				fmt.Println("All entries in the history file:")
				for _, ln := range results.All {
					fmt.Println(ln)
				}
			}
		},
	}
	showCmd.Flags().BoolVarP(&o.showBeginning, "beginning", "b", false, "Show the beginning entries (default 10).")
	showCmd.Flags().BoolVarP(&o.showLatest, "latest", "n", false, "Show the most recent entries (default 10).")
	showCmd.Flags().IntVarP(&o.showCount, "count", "c", 10, "The number of entries to show")
	showCmd.Flags().BoolVarP(&o.showDup, "dupes", "d", false, "Show the duplicate entries.")
	showCmd.Flags().BoolVarP(&o.showUniqueVals, "unique", "u", false, "Show unique entries.")
	showCmd.Flags().BoolVarP(&o.showMaxDup, "max-dup", "m", false, "(use with -d) Show the maximum duplicated entry")

	return showCmd
}

func init() {
	rootCmd.AddCommand(CreateShowCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
