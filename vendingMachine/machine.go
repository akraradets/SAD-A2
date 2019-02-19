package vendingMachine

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
func (m VendingMachine) display() {  
    fmt.Printf("Machine '%s'", m.Name)
}