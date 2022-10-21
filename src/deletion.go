package src

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	RE_REGEXP = `^((?:(?:[^?+*{}()[\]\\|]+|\\.|\[(?:\^?\\.|\^[^\\]|[^\\^])(?:[^\]\\]+|\\.)*\]|\((?:\?[:=!]|\?<[=!]|\?>)?(?1)??\)|\(\?(?:R|[+-]?\d+)\))(?:(?:[?+*]|\{\d+(?:,\d*)?\})[?+]?)?|\|)*)$`
)

func  (e *Engine)DeleteFolder(folder string) error { 
	err := filepath.WalkDir(folder, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		
		// if it is a file, delete it

		e.DeleteFile(path)
		return nil
	}) 

	if err != nil {
		return err
	}

	return os.Remove(folder)
}

func  (e *Engine)DeleteFile(file string)  error{
   // Open the file
   fd, err := os.OpenFile(file, os.O_RDWR, 0666)
   if err != nil {
	  return err
   }
   defer fd.Close()
   // Write random bytes to the file
   err = RandomiseFile(fd)

   if err != nil {
	  return err
   }

   return os.Remove(file)

}

func (e *Engine)UnlinkFile(file string) {

}

func (e *Engine)DeleteItem(item string) error {
	files, err := e.MatchWildCard(item)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errors.New("no files found")
	}

	for _, file := range files {
		// get the file descriptor
		fd, err := os.Stat(file)
		if err != nil {
			return err
		}

		// check if it is a directory
		if fd.IsDir() {
			if e.Input.HasFlag(FLAG_R) || e.Input.HasFlag(FLAG_D) {
				e.DeleteFolder(file)
			} else {
				return errors.New("cannot remove " + file + ": Is a directory")
			}
		} else {
			e.DeleteFile(file)
		}
	}

	return nil
}

func (e *Engine) MatchWildCard(wildcard string) ([]string,error) {
	return filepath.Glob(wildcard)
}