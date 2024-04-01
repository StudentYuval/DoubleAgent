package logger

/*
#include <android/log.h>

void logInfo(const char* msg) {
    __android_log_write(ANDROID_LOG_INFO, "GoLog", msg);
}
*/
import "C"
import "unsafe"

// LogInfo prints messages to Android's logcat at the info level.
func LogInfo(message string) {
	msg := C.CString(message)
	defer C.free(unsafe.Pointer(msg))
	C.logInfo(msg)
}
