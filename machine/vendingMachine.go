package machine

import (
    "fmt"
	"sync"
)

type vendingMachine struct {
    name   string
    count  int
    // LastName    string
    // TotalLeaves int
    // LeavesTaken int
}

var (
    once sync.Once
    instance *vendingMachine
)

// constructor
/* The suggested way from https://github.com/tmrts/go-patterns/blob/master/creational/singleton.md */
// func NewMachine(name string) vendingMachine {
// 	once.Do(func() {
// 		instance = vendingMachine{
//             Name: name, 
//             count: 0,
//         }
//     })
// 	return instance
// }
/* The suggested way from https://stackoverflow.com/questions/1823286/singleton-in-go */
func NewMachine(name string) *vendingMachine {
    once.Do(func() {
        instance = &vendingMachine{
            name: name,
            count: 0,
        }
    })
    return instance
}

func GetMachine() *vendingMachine{
    return instance
}


func (m *vendingMachine) Display() string {
	return fmt.Sprintf("Machine '%s'", m.name)
}

func (m *vendingMachine) Count() int {
    m.count = m.count + 1
    return m.count
}