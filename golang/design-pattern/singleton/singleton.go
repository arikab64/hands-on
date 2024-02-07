package singleton

import (
	"sync"
	"sync/atomic"
)

/**
 * This is a singleton example in Golang.
 * It uses sync.Once to ensure that only one instance is created.
 * It also uses sync/atomic to ensure that the count is incremented atomically.
 */

type Singleton interface {
	AddOne()
	GetCount() int64
}

type singleton struct {
	count int64
}

var instance *singleton
var once sync.Once

// GetInstance returns a singleton instance.
// It uses sync.Once to ensure that only one instance is created
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{count: 0}
	})
	return instance
}

// AddOne is a thread-safe method to increment the count of the singleton instance
// @see: https://golang.org/pkg/sync/atomic/
func (s *singleton) AddOne() {
	atomic.AddInt64(&s.count, 1)
}

// GetCount is a thread-safe method to get the count of the singleton instance
// It uses atomic.LoadInt64 to get the value of count atomically
func (s *singleton) GetCount() int64 {
	return atomic.LoadInt64(&s.count)
}
