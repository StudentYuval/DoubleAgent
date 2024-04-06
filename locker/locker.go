package locker

import (
	"fmt"
	"os"
	"syscall"
)

const perm = 0666

func Trylock(lockFile string) bool {
	file, err := os.OpenFile(lockFile, os.O_CREATE|os.O_RDWR, perm)

	if err != nil {
		fmt.Println("failed to open the file")
		return false
	}
	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB) // trying to lock
	if err != nil {
		fmt.Println("failed to lock the file - another instance is running!")
		file.Close()
		return false
	}
	// lock acquired
	fmt.Println("managed to acquire the lock - this is the 1st instance of the agent")
	return true
}
