# hist-man

## A simple history viewer for the bash_history file.

This project is a simple playground for a couple of things:
* get familiar with go and its conventions, patterns, etc.
* learn the spf13 cobra command library, conventions, and patterns.

### build/run
As of yet a build script has not been generated.  So the easiest way to run is with the "go run" command at the terminal.
go run hist-man.go [command]

The first couple of commands to be implemented with be simple:

### size
This command will give the size of the bash_history file.  The default will be in lines.  
The optional argument -b will give the size in bytes as well.
The optional argument -u will include the number of unique lines.  


### show
This command will simply print the entire contents of bash_history by default (same as the bash history command).  A number of arguments have been added to display subsets of the file:

	* -b		// show the beginning (oldest) 10 entries.
	* -n		// show the last (most recent) 10 entries.
	* -n -c 20	// show the last 20 entries in the file.
	* -u 		// show the unique entries.
	* -d		// show only the duplicated entries.
	* -dm		// show the single most duplicated entry.
