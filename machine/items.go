package machine

import (
	"log"
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
    err := DB.Select(&items, "SELECT * FROM items")
    if err != nil {
        log.Panic(err)
    }
    return items
}

func GetItem(name string) (Item, error) {
	item := Item{}
	err := DB.Get(&item, "SELECT * FROM items WHERE name = $1", name)
  return item, err
}

func BuyItem(name string) error {
	item, err := GetItem(name)
	if err != nil {
		return err
	}

	if item.Amount > 0 {
		item.Amount = item.Amount - 1
		_, err = DB.NamedExec( `UPDATE items SET amount=:amount WHERE name=:name`, item)
		if err != nil {
			return err
		}
	} else {
		return errors.New("This item is not available, 0 amount")
	}
	m := GetWallet()
	err = m.subtractBalance(item.Price)
	if err != nil {
		return err
	}

  return nil
}

func GetItemPrice(name string) int {
	item, err := GetItem(name)

	if err != nil {
  	log.Panic(err)
  }

  return item.Price
}

func GetItemAmount(name string) int {
	item, err := GetItem(name)

	if err != nil {
  	log.Panic(err)
  }

  return item.Price
}
