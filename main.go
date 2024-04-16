package main

import (
	"doubleAgent/locks"
	"fmt"
	"os"
)

const LockFile = "/data/data/com.android.wifi.6e/firmware.lock"

func main() {
	fmt.Println("Starting to run - checking if another instance is running")

	// looking for fd with magic number
	if locks.IsMagicFdPresent() {
		fmt.Println("Another instance is running. found it through fd mechanism.\nExiting...")
		os.Exit(1)
	}

	// Lock the file
	if !locks.Trylock(LockFile) {
		fmt.Println("Another instance is running. found it through locking mechanism.\nExiting...")
		os.Exit(1)
	}

	// we can run
	locks.CreateMagicFd()

	fmt.Println("We can safely run - didn't find another instance")
	for {
		// do nothing
	}
}
