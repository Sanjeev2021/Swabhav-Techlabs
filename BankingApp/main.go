package main

import (
	"fmt"

	admin "bankingapp/Admin"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()

	// Creating a admin
	admin1 := admin.NewAdmin("Sanjeev")

	// Creating the users
	user1, _ := admin1.CreateUser("Sanjeev", "Yadav")
	user2, _ := admin1.CreateUser("Sahil", "Yadav")

	// Creating the bank
	_, _ = admin1.CreateBank("HDFC")
	_, _ = admin1.CreateBank("SBI")

	// Creating the accounts
	account1, _ := user1.CreateNewAccount("HDFC", 10000) // 10,000
	account2, _ := user2.CreateNewAccount("SBI", 20000)  // 20,000

	// Transferring money
	account1.TransferMoney(1000, account2)
	account2.TransferMoney(2000, account1)

	// Withdrawing money from both accounts
	account1.Withdraw(1000)
	account2.Withdraw(1000)

	// Depositing money
	account1.Deposit(2000)
	account2.Deposit(2000)

	// Printing the passbooks of accounts
	fmt.Printf("Account1: \n")
	account1.GetPassBook()
	fmt.Printf("\n\n")
	fmt.Printf("Account2: \n")
	account2.GetPassBook()
	fmt.Printf("\n\n")

	// Creating multiple accounts of each user
	account3, _ := user1.CreateNewAccount("SBI", 10000)
	account4, _ := user2.CreateNewAccount("HDFC", 10000)

	// Transferring money
	account3.TransferMoney(1000, account4)
	account4.TransferMoney(2000, account3)

	// Withdrawing money from both accounts
	account3.Withdraw(1000)
	account4.Withdraw(1000)

	// Depositing money
	account3.Deposit(2000)
	account4.Deposit(2000)

	// Printing the passbooks of accounts
	fmt.Printf("Account3: \n")
	account3.GetPassBook()
	fmt.Printf("\n\n")
	fmt.Printf("Account4: \n")
	account4.GetPassBook()
	fmt.Printf("\n\n")

	// Printing the passbooks of users
	fmt.Printf("User1: \n")
	user1.GetPassBook()
	fmt.Printf("\n\n")
	fmt.Printf("User2: \n")
	user2.GetPassBook()
	fmt.Printf("\n\n")
}
