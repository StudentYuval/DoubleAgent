package locks

// will create fd associated with this process - and store magic number in it
// this will be used to check if another instance is running
//

// check: itterate over all running processes, and check if the magic number is in the fd

// if it is - return true

// if not - return false

import (
	"fmt"
	"os"
)

const fd_path = "/proc/%d/fd/%d"
const magic_number = 0x12345678

func IsMagicFdPresent() bool {

	// iterate over all running processes, and check if the magic number is in the fd

	psList, err := os.ReadDir("/proc")
	if err != nil {
		fmt.Println("failed to read /proc")
		panic(err)
	}

	for _, ps := range psList {
		if ps.IsDir() {
			continue
		}
		pid := ps.Name()

		//create list of all fds of current process
		fdList, err := os.ReadDir(fmt.Sprintf("/proc/%d/fd", pid))
		if err != nil {
			fmt.Println("failed to read the fd directory")
			panic(err)
		}

		// iterate over all fds
		for fd := range fdList {

			// check if the file is a directory
			_, err := os.Stat(fmt.Sprintf(fd_path, pid, fd))
			if err != nil { // it is a directory
				continue
			}

			// check if the magic number is in the fd
			fdFile, err := os.Open(fmt.Sprintf(fd_path, pid, fd))
			if err != nil {
				fmt.Println("failed to open the fd file")
				continue
			}

			var magic int
			_, err = fmt.Fscanf(fdFile, "%d", &magic)
			if err != nil {
				fmt.Println("failed to attempt reading the magic number")
				continue
			}
			if magic == magic_number {
				return true
			}
		}
	}
	return false
}
