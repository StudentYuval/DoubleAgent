package main

import (
	"doubleAgent/ipc"
	"doubleAgent/locker"
	"fmt"
)

const LockFile = "/data/data/com.android.wifi.6e/firmware.lock"

func main() {
	fmt.Println("Starting to run - checking if another instance is running")

	//// Check if another instance is running
	//if !ipc.CheckInstance() {
	//	return
	//}

	// Lock the file
	if !locker.Trylock(LockFile) {
		fmt.Println("Another instance is running. found it through locking mechanism.\nExiting...")
		return
	}

	// we can run - start the listener
	if err := ipc.InitListener(); err != nil {
		fmt.Println("Error starting listener:", err)
		return
	}

	fmt.Println("Listener started. Waiting for connections...")
	for {
	}
}
