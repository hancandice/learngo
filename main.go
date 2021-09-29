package main

import (
	"fmt"

	"github.com/hancandice/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Jeeyoung Han")
	account.Deposit(10000)
	fmt.Println(account)
}