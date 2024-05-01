package locks

/*
#include <sys/resource.h>

int set_soft_limit(int new_limit){
	struct rlimit limit;

	// make sure current limit is not NEW_LIMIT
	if (getrlimit(RLIMIT_NOFILE, &limit) == -1){
		return -1;
	}

	if (limit.rlim_cur == NEW_LIMIT){
		return 0; // already set
	}

	// set the new limit
	limit.rlim_cur = NEW_LIMIT;
	if (setrlimit(RLIMIT_NOFILE, &limit) == -1){
		return -1;
	}

	return 1; // success
}
*/
import "C"
import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const MAGIC_LIMIT = 1234
const DEFAULT_ANDROID_LIMIT = 2 << 9 // TODO: make sure this is correct

type LimitLock struct{}
type LimitLockFactory struct{}

func (f *LimitLockFactory) NewLock() Lock {
	return &LimitLock{}
}

const pattern = "Max open files.*1234"
const limitsPath = "/proc/%d/limits"

func (l *LimitLock) Acquire() error {
	// set the new soft limit
	res := int(C.set_soft_limit(MAGIC_LIMIT))

	switch res {

	case -1:
		return fmt.Errorf("failed to set new soft limit")
	case 0:
		return fmt.Errorf("soft limit is already set")
	case 1:
		return nil // success
	default:
		return fmt.Errorf("unexpected return value from C.set_soft_limit")

	}
}
func (l *LimitLock) IsLocked() bool {
	// search for the soft limit in existing /proc files
	psList, err := os.ReadDir("/proc")
	if err != nil {
		panic(err)
	}

	for _, ps := range psList {
		if _, err := strconv.Atoi(ps.Name()); err == nil { // folder name is a PID
			limitsFile, err := os.Open(fmt.Sprintf(limitsPath, ps.Name()))
			if err != nil {
				continue
			}

			// check if the file contains the soft limit
			content := make([]byte, 1024)
			_, err = limitsFile.Read(content)
			if err != nil {
				continue
			}

			// Define the regex pattern
			re := regexp.MustCompile(pattern)
			if re.Match(content) {
				return true
			}
		}
	}
	return false
}

func (l *LimitLock) Release() error {
	// set the default android soft limit
	res := int(C.set_soft_limit(DEFAULT_ANDROID_LIMIT))

	switch res {

	case -1:
		return fmt.Errorf("failed to set new soft limit")
	case 0:
		return fmt.Errorf("default soft limit is already set - nothing to release")
	case 1:
		return nil // success
	default:
		return fmt.Errorf("unexpected return value from C.set_soft_limit")
	}
}
