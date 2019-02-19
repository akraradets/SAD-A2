package singleton

import "sync"

type singleton struct{
	Name string
}

var (
    once sync.Once

    instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = singleton{Name:"ASDASD"}
	})

	return instance
}