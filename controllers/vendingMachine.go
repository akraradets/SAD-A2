package controllers

import (
	"github.com/SAD-A2/machine"
    // "fmt"
)

type VendingMachine struct {
	Machine   machine.VendingMachine
}

// GET /vendingMachine/Index
func (controller VendingMachine) Index() string{
	return "Your index"
}
// GET /vendingMachine/Name
func (controller VendingMachine) Name() string{
	result := controller.Machine.Display()
	return result
}