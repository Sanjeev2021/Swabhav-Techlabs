package main

import (
	"fmt"

	//"bankingapp/Account"
	admin "bankingapp/Admin"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	admin1 := admin.NewAdmin("SANJEEV")
	user1, _ := admin1.CreateUser("Sanjeev", "Yadav")

	//admin1.DeleteCreatedUser("Sanjeev")

	//fmt.Println(admin1)

	bank1, _ := admin1.CreateBank("SBI")
	fmt.Println(bank1)

	// _, err := user1.CreateNewAccount("HDFC", 6883738, "SADSFWF12")
	// if err != nil {
	// 	panic(err)
	// }

	account1, _ := user1.CreateNewAccount("SBI", 10000000)

	err := admin1.UpdateCreatedUser("Sanjeev", "firstName", "SAHIL")
	if err != nil {
		panic(err)
	}

	//admin1.UpdateCreatedBank("HDFC", "bankname", "SBI")

	//admin1.DeleteCreatedBank("HDFC")
	//fmt.Println(admin1.BanksCreatedByMe)
	//fmt.Println(admin1.BanksCreatedByMe)
	// CHECK BANK UPDATE
	//admin1.UpdateCreatedBank("SBI", "bankname", "HDFC")
	//fmt.Println(admin1.BanksCreatedByMe[0])

	account2, err := user1.CreateNewAccount("SBI", 1000)
	if err != nil {
		fmt.Println("err: ", err)
	}
	account1.TransferMoney(10000, account2)
	account2.TransferMoney(2000, account1)
	account1.Withdraw(100)
	user1.GetPassBook()
}
