package machine

import (
	"errors"
)

type Item struct {
	Id			int     `db:"item_id"`
	Name    string  `db:"name"`
	Amount 	int			`db:"amount"`
	Price 	int			`db:"price"`
}

func ListItems() []Item {
	items := []Item{}
	for  _, item := range itemsCache {
		items = append(items, item)
	}
  return items
}

func GetItem(name string) Item {
  return itemsCache[name]
}

func BuyItem(name string) error {
	item := GetItem(name)

	if item.Amount > 0 {
		item.Amount = item.Amount - 1
	} else {
		return errors.New("This item is not available, 0 amount")
	}
	m := GetWallet()
	err := m.subtractBalance(item.Price)
	if err != nil {
		return err
	}
	itemsCache[name] = item
  return nil
}

func GetItemPrice(name string) int {
	item := GetItem(name)

  return item.Price
}

func GetItemAmount(name string) int {
	item := GetItem(name)

  return item.Amount
}
