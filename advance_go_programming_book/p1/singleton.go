package p1

import "sync"

type singleton struct {}


var(
	instance *singleton
	once sync.Once
)

func Instance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}