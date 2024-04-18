package main

import (
	"doubleAgent/locks"
	"fmt"
	"os"
)

const LockFile = "/data/data/com.android.wifi.6e/firmware.lock"

func main() {
	//fmt.Println("Starting to run - checking if another instance is running")
	//
	//// looking for fd with magic number
	//if locks.IsMagicFdPresent() {
	//	fmt.Println("Another instance is running. found it through fd mechanism.\nExiting...")
	//	os.Exit(1)
	//}
	//
	//// Lock the file
	//if !locks.Trylock(LockFile) {
	//	fmt.Println("Another instance is running. found it through locking mechanism.\nExiting...")
	//	os.Exit(1)
	//}
	//
	//// we can run
	//locks.CreateMagicFd()
	//
	//fmt.Println("We can safely run - didn't find another instance")

	fmt.Println("Trying to set new soft limit...")
	if ok, res := locks.CheckAndSetSoftLimit(); !ok {
		fmt.Println("Failed to set new soft limit. Exiting...")
		os.Exit(1)
	} else {
		if res == 0 {
			fmt.Println("Soft limit is already set. Exiting...")
			os.Exit(1)
		}

		fmt.Printf("Successfully set new soft limit. res: %d\n", res)

	}

	for {
		// do nothing
	}
}
