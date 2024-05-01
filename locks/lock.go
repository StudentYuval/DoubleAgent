package locks

type Lock interface {
	Acquire() error
	Release() error
	IsLocked() bool
}

type LockFactory interface {
	NewLock() Lock
}

func GetLock(lockType string) Lock {
	switch lockType {
	case "fd":
		factory := &FdLockFactory{}
		return factory.NewLock()
	case "limit":
		factory := &LimitLockFactory{}
		return factory.NewLock()
	default:
		return nil
	}
}
