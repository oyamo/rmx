package src
import(
	"fmt"
	"os"
)

// PrintErr prints the error message to stderr.
func PrintErr (err string) {
	fmt.Fprintf(os.Stderr, "rmx: %s", err)
	fmt.Fprintln(os.Stderr)
}

// PrintErrf prints the error message to stderr.
func PrintErrf (format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "rmx: " + format + "", args...)
	fmt.Fprintln(os.Stderr)
}

// PrintErr prints the info message to stdout.
func PrintInfof(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format, args...)
	fmt.Fprintln(os.Stdout)
}

// PrintErr prints info message to stdout.
func PrintInfo (info string) {
	fmt.Fprint(os.Stdout, info)
	fmt.Fprintln(os.Stdout)
}

