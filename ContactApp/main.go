package main

import (
	"fmt"

	"contactApp/user"

)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered From", r)
		}
		fmt.Println("End of Main")
	}()
	admin1 := user.NewAdmin("Yash", "Shah", "yashshah")
	fmt.Println(admin1)

	user1, err := admin1.NewUser("Ankush", "Sondal", "ankushsondal")
	if err != nil {
		panic(err)
	}
	// fmt.Println(user1)
	// user2, err := admin1.NewUser("Sanjeev", "Yadav", "sanjeev21")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(user2)

	// fmt.Println(admin1.ReadNewUser())
	// fmt.Println(admin1.UpdatedUser("sanjeev21", "firstName", "sahil"))
	// fmt.Println(admin1.DeleteUser("sanjeev21"))
	// yash, _ := admin1.NewUser("yash", "shah", "yashshah")
	// yash.CreateContactUser("Sanjeev")
	// //yash.
	// //fmt.Println(user1.CreateContactUser("sanjeev"))
	// //fmt.Println(user1.UpdateContactUser("sanjeev", "sahil"))
	// //fmt.Println(user1.DeleteContactUser("sahil"))
	// //fmt.Println(user1.FindContactInfo())
	// fmt.Println("CREATE CONTACTINFO>>>>>>>>>")
	// yash.CreateContactInfo("Sanjeev", "mob", "123455")
	// fmt.Println(yash.Mycontacts[0])
	// //yash.DeleteContactInfo("fhfhytdt","mob")
	// fmt.Println("UPDATE CONTACTINFO>>>>>>>>>>>>>>>>>>>>>>>>")
	// yash.UpdateContactInfo("Sanjeev", "email", "sanjeev@gmail")
	// fmt.Println(yash.Mycontacts[0].MyContactInfo)

	// fmt.Println("DELETE CONTACTINFO >>>>>>>>>>>>>>>>>>")
	// yash.DeleteContactInfo("Sanjeev", "mob")
	// fmt.Println(yash.Mycontacts[0])
	// // give details of all the users created by admin

	user1.CreateContactUser("ankush")
	fmt.Println(user1)
	fmt.Println("contact >>>>>>>", user1.Mycontacts[0])

	user1.CreateContactInfo("ankush", "mob", "123456")
	fmt.Println("contact info >>>>>>>", user1.Mycontacts[0].MyContactInfo[0])

	user1.UpdateContactInfo("ankush", "mob", "999999")
	fmt.Println("contact info >>>>>>>", user1.Mycontacts[0].MyContactInfo[0].ContactInfoValue)

	user1.DeleteContactInfo("ankush", "mob")
	fmt.Println("contact >>>>>>>", user1.Mycontacts[0].MyContactInfo)

}
