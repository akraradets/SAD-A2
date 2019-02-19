package machine

import (
    "fmt"
)

type VendingMachine struct {
    Name   string
    // LastName    string
    // TotalLeaves int
    // LeavesTaken int
}

// constructor
func (m VendingMachine) Display() string {
	return fmt.Sprintf("Machine '%s'", m.Name)
}
