/**
 * Writes pseudo-random bytes into a file or zeroes it.
 *
**/

package src

import (
	"bufio"
	"crypto/rand"
	"io"
	"os"
)

const (
	BUFFER_SIZE = 1024
)

func ZeroFile(file string) {
 
}


func RandomiseFile(file *os.File) error {
	// use io.CopyN to write random bytes to the file
	fb := bufio.NewWriter(file)
	defer fb.Flush()

	// Get the block size of the file
	stat, err := file.Stat()
	if err != nil {
	   return err
	 }
 
	 // 
	size := stat.Size()

	_, err = io.CopyN(fb, rand.Reader, size)

	return err
}