package src

import (
	"flag"
	"fmt"
	"os"
)

const (
	FLAG_F = iota  // -f, --force
	FLAG_I // -i
	FLAG_R // -r, -R, --recursive
	FLAG_D // -d, --dir
	FLAG_V // -v, --verbose
	FLAG_HELP // --help, -h
)

type Input struct {
	Files []string
	Flags []uint8 // tiny for less mem
}


type InputParser interface {
	HasFlag(uint8) bool
}


// NewInput creates a new Input struct and parses the command line arguments.
// It returns the Input struct and an error if there is one.
func NewInput() (*Input, error) {
	args := os.Args[1:]

	if len(args) == 0 {
		return nil, fmt.Errorf("no arguments")
	}

	input := &Input{
		Files: make([]string, 0),
		Flags: make([]uint8, 0),
	}

	// set flag usage
	flag.Usage = Usage

	var help = flag.Bool("help", false, "")
	var version = flag.Bool("version", false, "")
	var force = flag.Bool("f", false, "")
	var forceShort = flag.Bool("force", false, "")
	var recursive = flag.Bool("recursive", false, "")
	var recursiveShort = flag.Bool("r", false, "")
	var dir = flag.Bool("dir", false, "dir")
	var dirShort = flag.Bool("d", false, "dir")
	var verbose = flag.Bool("verbose", false, "")
	var verboseShort = flag.Bool("v", false, "")
	var i = flag.Bool("i", false, "i")

	flag.Parse()

	if *help {
		input.Flags = append(input.Flags, FLAG_HELP)
		return input, nil
	}

	if *version {
		input.Flags = append(input.Flags, FLAG_V)
		return input, nil
	}

	if *force  || *forceShort {
		input.Flags = append(input.Flags, FLAG_F)
	}

	if *recursive || *recursiveShort {
		input.Flags = append(input.Flags, FLAG_R)
	}

	if *dir || *dirShort {
		input.Flags = append(input.Flags, FLAG_D)
	}

	if *verbose || *verboseShort {
		input.Flags = append(input.Flags, FLAG_V)
	}

	if *i {
		input.Flags = append(input.Flags, FLAG_I)
	}

	input.Files = flag.Args()

	return input, nil
}

// HasFlag returns true if the input has the flag.
func (i *Input) HasFlag(flag uint8) bool {
	for _, f := range i.Flags {
		if f == flag {
			return true
		}
	}
	return false
}