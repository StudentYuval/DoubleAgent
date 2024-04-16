package main

import (
	"doubleAgent/locks"
	"fmt"
)

const LockFile = "/data/data/com.android.wifi.6e/firmware.lock"

func main() {
	fmt.Println("Starting to run - checking if another instance is running")

	// looking for fd with magic number
	if locks.IsMagicFdPresent() {
		fmt.Println("Another instance is running. found it through fd mechanism.\nExiting...")
		return
	}

	// Lock the file
	if !locks.Trylock(LockFile) {
		fmt.Println("Another instance is running. found it through locking mechanism.\nExiting...")
		return
	}

	// we can run - start the listener

	fmt.Println("We can safely run - didn't find another instance")
	for {
		// do nothing
	}
}
