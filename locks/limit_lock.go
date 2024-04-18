package locks

/*
#include <sys/resource.h>
#include <stdio.h>

#define NEW_LIMIT 1234

int set_soft_limit(){
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

func CheckAndSetSoftLimit() (bool, int) {
	res := int(C.set_soft_limit())

	switch res {

	case -1:
		return false, res
	case 0:
		return true, res
	case 1:
		return true, res
	default:
		return false, res
	}

}
