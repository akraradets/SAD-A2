package machine

import (
    "fmt"
	"sync"
)

type VendingMachine struct {
    name   string
    count  int
    // LastName    string
    // TotalLeaves int
    // LeavesTaken int
}

var (
    once sync.Once
    instance VendingMachine
)

// constructor
func New(name string) VendingMachine {
	once.Do(func() {
		instance = VendingMachine{
            name: name, 
            count: 0
        }
    })
	return instance
}

func (m VendingMachine) Display() string {
	return fmt.Sprintf("Machine '%s'", m.Name)
}

func (m VendingMachine) count() string {
    m.count = m.count + 1
    return m.count
}