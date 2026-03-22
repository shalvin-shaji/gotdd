package pointerr

import "fmt"
import "errors"

type Bitcoin int

type Wallet struct {
	Value Bitcoin
}

var ErrInsuffientFunds = errors.New("Insufficient funds")

func (w *Wallet) Deposit(money Bitcoin) {
	fmt.Printf("address of balance in deposit is %p \n", &w.Value)
	w.Value += money
}

func (w *Wallet) Balance() Bitcoin {
	return w.Value
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (b *Wallet) Withdraw(amount Bitcoin) error {
	if b.Value < amount {
		return ErrInsuffientFunds
	}
	b.Value -= amount
	return nil

}
