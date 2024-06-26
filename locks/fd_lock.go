package locks

import (
	"fmt"
	"os"
)

// TODO: change this logic to use mem_fd

type FdLock struct{}
type FdLockFactory struct{}

func (f *FdLockFactory) NewLock() Lock {
	return &FdLock{}
}

const fd_path = "/proc/%d/fd/%d"
const magic_number = 0x12345678

func (l *FdLock) IsLocked() bool {

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

func (l *FdLock) Acquire() error {
	// get my pid
	pid := os.Getpid()

	// create a file with the magic number
	fdFile, err := os.OpenFile(fmt.Sprintf(fd_path, pid, 3), os.O_CREATE|os.O_RDWR, 0666)
	defer fdFile.Close()
	if err != nil {
		fmt.Println("failed to open the fd file")
		return err
	}

	_, err = fmt.Fprintf(fdFile, "%d", magic_number)
	if err != nil {
		fmt.Println("failed to write the magic number")
		return err
	}
	fmt.Println("created the magic fd")
	return nil
}

func (l *FdLock) Release() error {
	// get my pid
	pid := os.Getpid()

	// remove the file with the magic number
	err := os.Remove(fmt.Sprintf(fd_path, pid, 3))
	if err != nil {
		fmt.Println("failed to remove the fd file")
		return err
	}
	fmt.Println("removed the magic fd")
	return nil
}
