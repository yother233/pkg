package mutex

import (
	"errors"
	"sync"
)

var (
	ErrNotAcquireLock = errors.New("failed t acquire lock")
)

type Locker interface {
	Lock(key string) error
	Unlock(key string) error
}

type DefaultLock struct {
	mu sync.Mutex
}

// Lock implements Locker.
func (d *DefaultLock) Lock(_ string) error {
	d.mu.Lock()
	return nil
}

// Unlock implements Locker.
func (d *DefaultLock) Unlock(key string) error {
	d.mu.Unlock()
	return nil
}
