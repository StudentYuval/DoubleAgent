package DoubleAgent

import (
	"doubleAgent/logger"
	"runtime"
)

func main() {

	logger.LogInfo("Hello Logcat!!")
	// Crash Detector and Semaphore Initialization:

	// 1. Crash detector and semaphore are initialized using their respective
	//    constructors (`NewCrashDetector` and `NewSemaphore`). These functions
	//    handle resource creation (temporary file, memory mapping, semaphore)
	//    and potential errors. `defer` statements ensure proper cleanup (`Close` methods)
	//    even on program exit.

	//TODO: implement

	// Semaphore Acquisition:

	// 2. Attempt to acquire the semaphore using `sem.Acquire()`. This
	//    serializes access to the critical section, ensuring only one instance
	//    can proceed at a time. If another instance holds the semaphore
	//    (indicating it's running), the current instance might block until it's
	//    released. Errors during acquisition are captured.

	//TODO: implement

	// Defer releasing the semaphore in a critical section. This ensures it's
	// released even if the program crashes within the critical section,
	// preventing indefinite blocking for other instances.

	// Critical Section:

	// 3. Check if this is the first instance using `detector.IsFirstInstance()`.
	//    This method (in crash_detector.go) examines the shared state in the
	//    memory-mapped file and updates it atomically if necessary.

	// Start Periodic Updates (Optional):

	// 4. If this is the first instance, launch a goroutine using
	//    `go detector.RunPeriodicUpdates()`. This goroutine (defined in
	//    crash_detector.go) periodically updates the shared state to indicate
	//    the program is still running, improving crash detection.

	//TODO: implement on 2nd phase

	// Your program logic here

	// After your program logic completes (or encounters an error), execution
	// continues here. The deferred `sem.Release()` ensures the semaphore is
	// released, allowing other potential instances to acquire it and proceed.

	runtime.Goexit() // Exit the program gracefully
}
