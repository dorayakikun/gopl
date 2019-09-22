package bank

var deposits = make(chan int)
var balances = make(chan int)
var withdrawals = make(chan int)
var result = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	withdrawals <- amount
	return <-result
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdrawals:
			ok := amount < balance
			if ok {
				balance -= amount
			}
			result <- ok
		}
	}
}

func init() {
	go teller()
}
