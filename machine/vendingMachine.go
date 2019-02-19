package machine

import (
    "fmt"
	"sync"
)

type VendingMachine struct {
    Name   string
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
func NewMachine(name string) VendingMachine {
	once.Do(func() {
		instance = VendingMachine{
            Name: name, 
            count: 0,
        }
    })
	return instance
}

func (m VendingMachine) Display() string {
	return fmt.Sprintf("Machine '%s'", m.Name)
}

func (m VendingMachine) Count() int {
    m.count = m.count + 1
    return m.count
}