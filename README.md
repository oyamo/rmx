# rmx : Military Grade Deletion
Usually when you delete a file using `rm`, It removes the pointer to the file system. However, the data can still be accessed by forensic/data recovery tools.

`rmx` Shreds your files/folders by overwriting them with multiple rounds of random data. It also removes the pointer to the file system. This makes it impossible to recover the data.

## Installation
```bash
go install github.com/oyamo/rmx
echo "export PATH=$PATH:/home/$USER/go/bin" >> ~/.bashrc
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
- [ ] Zero out the file
- [ ] Randomize the file name
- [ ] Update makefile to install the binary in /usr/local/bin
- [ ] Add a progress bar
- [ ] Create deb and rpm packages

