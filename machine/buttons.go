package machine

import (
	"errors"
)

// To use proxy and to object they must implement same methods
type ButtonInterface interface {
    Push() error
}

// Button represents real objects which proxy will delegate data
type button struct {
    name string
}

// Button implements IObject interface and handel's all logic
func (but *button) Push() error {
    // Action behavior
    err := BuyItem(but.name)
    return err
}

// ProxyObject represents proxy object with intercepts actions
type ProxyButton struct {
    button button
}

// ProxyButton are implemented ButtonInterface and intercept action before send in real Object
func (p *ProxyButton) Push() error {
    m := GetWallet();
    name := p.button.name
    currentBalance := m.CheckBalance()
    price := GetItemPrice(name)
    if price <= currentBalance {
      return errors.New("Not enough balance to buy the item.")
    }
    err := p.button.Push()
    return err
}

// ProxyButton builder
func NewProxyButton(name string) *ProxyButton {
    but := new(ProxyButton)
    but.button = button{name: name}
    return but
}
