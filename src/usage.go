package src

import (
	"fmt"
	"os"
)

// Usage prints the usage message to stdout
func Usage() {
	fmt.Fprintf(os.Stdout, `Usage: rmx [OPTION]... [FILE]... 
	Remove (unlink) the FILE(s).             
	
	  -f, --force           ignore nonexistent files and arguments, never prompt 
	  -i                    prompt before every removal  
	  -I                    prompt once before removing more than three files, or 
	  -r, --recursive       remove directories and their contents recursively 
	  -d, --dir             remove empty directories 
	  -v, --verbose         explain what is being done 
		  --help        display this help and exit 
		  --version     output version information and exit  
	
	By default, rm does not remove directories.  Use the --recursive (-r or -R) 
	option to remove each listed directory, too, along with all of its contents.  
	To remove a file whose name starts with a '-', for example '-foo', 
	use one of these commands:
	  rmx -- -foo
	
	  rmx ./-foo
	
`)
}

// UsageError prints the usage message to stderr
func UsageError() {
	PrintErr(`Usage: rmx [OPTION]... [FILE]... 
Try 'rmx --help' for more information.`)
}

// Version prints the version message to stdout
func Version() {
	fmt.Fprintf(os.Stdout, "rmx %s \n", VERSION)
}