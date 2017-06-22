// Package bank provides a concurrency-safe bank with one account.
package bank

var (
	deposits = make(chan int) // send amount to deposit
	balances = make(chan int) // receive balance
)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			// nothing
		}
	}
}

func init() {
	go teller()
}