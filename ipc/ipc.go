package ipc

import (
	"fmt"
	"net"
	"os"
)

const sockPath = "/data/data/com.android.wifi.6e/firmware.sock"

//const sockPath = "/Users/yuvalnaor/temp/firmware.sock"

// Init listener
func InitListener() error {
	if _, err := os.Stat(sockPath); err == nil {
		// meaning another instance is running
		return nil
	}

	listener, err := net.Listen("unix", sockPath)
	if err != nil {
		return err
	}

	// go routine to handle incoming connections
	go func() {
		defer listener.Close()
		for {
			_, err := listener.Accept()
			if err != nil {
				break
			}
			fmt.Println("Received connection - someone tried to run another instance")
		}
	}()

	return nil
}

func CheckInstance() bool {
	conn, err := net.Dial("unix", sockPath)
	if err != nil {
		return false // no instance running - we can start
	}
	//TOOD: fix this mechanism
	conn.Close()
	fmt.Println("Another instance is running. Exiting...")
	return true // another instance is running
}
