package machine
import (
	
	"log"
	"strconv"
)
type Item struct {
	Id		int     `db:"item_id"` 
	Name    string  `db:"name"`
	Amount 	int		`db:"amount"`
	Price 	int		`db:"price"`
}

func ListItems() []Item {
	items := []Item{}
    err := DB.Select(&items, "SELECT * FROM items")
    if err != nil {
        log.Panic(err)
    }
    return items
}

func BuyItems(name string) string {
	item := Item{}
	err := DB.Get(&item, "SELECT * FROM items WHERE name = $1", name)
	// amount_str := strconv.Itoa(item.amount)
	if err == nil && item.Amount > 0 {
		item.Amount = item.Amount - 1 
		_, err = DB.NamedExec( `UPDATE items SET amount=:amount WHERE name=:name`, item)
	}

	if err != nil {
        log.Panic(err)
    }
    return strconv.Itoa(item.Amount)
} 