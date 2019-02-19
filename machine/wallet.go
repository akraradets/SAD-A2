package machine

import (
    "fmt"
	"sync"
)

type wallet struct {
    name   string
    balance  int
    // LastName    string
    // TotalLeaves int
    // LeavesTaken int
}

var (
    once sync.Once
    instance *wallet
)

// constructor
/* The suggested way from https://github.com/tmrts/go-patterns/blob/master/creational/singleton.md */
// func NewWallet(name string) wallet {
// 	once.Do(func() {
// 		instance = wallet{
//             Name: name, 
//             balance: 0,
//         }
//     })
// 	return instance
// }
/* The suggested way from https://stackoverflow.com/questions/1823286/singleton-in-go */
func NewWallet(name string) *wallet {
    once.Do(func() {
        instance = &wallet{
            name: name,
            balance: 0,
        }
    })
    return instance
}

func GetWallet() *wallet{
    return instance
}


func (m *wallet) Display() string {
	return fmt.Sprintf("Wallet name is '%s'", m.name)
}

/* checkBalance */
func (m *wallet) CheckBalance() int {
    return m.balance
}
/* InsertCoin */
func (m *wallet) InsertCoin(coin int) {
    m.balance = m.balance + coin
}
/* RetriveCoin */
func (m *wallet) RetriveCoin() {
    m.balance = 0
}