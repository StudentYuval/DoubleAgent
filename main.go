package main

import (
	"doubleAgent/locks"
	"doubleAgent/logic"
)

func main() {
	// getting a lock from the factory and locking it
	lock := locks.GetLock("limit")
	if lock == nil {
		panic("Lock not found")
	}

	// check if the lock is acquired - if not, acquire it
	// if it does - try fd lock.

	if !lock.IsLocked() {
		err := lock.Acquire()
		if err != nil {
			panic(err)
		}
	} else {
		lock = locks.GetLock("fd")
		if lock == nil {
			panic("Lock not found")
		}
		if !lock.IsLocked() {
			err := lock.Acquire()
			if err != nil {
				panic(err)
			}
		}
	}

	go logic.StartServer()

	select {}

	err := lock.Release()
	if err != nil {
		panic(err)
	}

}
