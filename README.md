# rmx : Military Grade Deletion
rmx is a command line tool for securely deleting files and folders. 
Depending on the file size, `rmx` will take a long time to delete the file than `rm`.

## Installation
```bash
go get github.com/oyamo/rmx

```

## Screenshots
![rmx](assets/Screenshot%20from%202022-10-21%2016-45-05.png)

## Progress
- [x] Delete files
- [x] Delete folders
- [x] Delete files in a folder
- [x] Delete files in a folder recursively
- [x] Delete files in a folder recursively with a regex pattern
- [x] Write random bytes to the file
- [] Zero out the file
- [] Randomize the file name
- [] Update makefile to install the binary in /usr/local/bin
- [] Add a progress bar
- [] Create deb and rpm packages

