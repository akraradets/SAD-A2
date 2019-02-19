package controllers

import (
	// "github.com/SAD-A2/machine"
    // "fmt"
)

type VendingMachine struct {
	Name: "CSIM Vending Machine"
}

// Constructor
func NewVendingMachine() *vendingMachine {
    m := new(vendingMachine)
    return m
}

// GET /VendingMachine/Index
func (controller vendingMachine) Index() VendingMachine{
	data{balance: machine.wallet.CheckBalance()}
	return data
}