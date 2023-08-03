package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"bankingApp/controller"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()

	HandleFunction()
}

func HandleFunction() {
	// Define allowed headers, origins, and methods for CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router := mux.NewRouter().StrictSlash(true)

	router = router.PathPrefix("/api/v1/bankingApp").Subrouter()

	//Routes for user management
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", controller.CreateUser).Methods("POST")
	userRouter.HandleFunc("/get/{id}", controller.GetUserById).Methods("GET")
	userRouter.HandleFunc("/getall/{page}", controller.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/update/{id}", controller.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/delete/{id}", controller.DeleteUser).Methods("DELETE")

	//Routes for Admin management
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/", controller.CreateAdmin).Methods("POST")
	adminRouter.HandleFunc("/get/{id}", controller.GetAdmin).Methods("GET")
	adminRouter.HandleFunc("/getall/{page}", controller.GetAllAdmins).Methods("GET")
	adminRouter.HandleFunc("/update/{id}", controller.UpdateAdmin).Methods("PUT")
	adminRouter.HandleFunc("/delete/{id}", controller.DeleteAdmin).Methods("DELETE")

	//Routes for bank management
	bankRouter := router.PathPrefix("/{adminid}/bank").Subrouter()
	bankRouter.HandleFunc("/", controller.CreateBank).Methods("POST")
	bankRouter.HandleFunc("/get/{id}", controller.GetBank).Methods("GET")
	bankRouter.HandleFunc("/getall/{page}", controller.GetAllBanks).Methods("GET")
	bankRouter.HandleFunc("/update/{id}", controller.UpdateBank).Methods("PUT")
	bankRouter.HandleFunc("/delete/{id}", controller.DeleteBank).Methods("DELETE")

	//Routes for account management
	accountRouter := router.PathPrefix("/{userid}/account").Subrouter()
	accountRouter.HandleFunc("/", controller.CreateAccount).Methods("POST")
	accountRouter.HandleFunc("/get/{id}", controller.GetAccount).Methods("GET")
	accountRouter.HandleFunc("/getall", controller.GetAllAccounts).Methods("GET")
	accountRouter.HandleFunc("/update/{id}", controller.UpdateAccount).Methods("PUT")
	accountRouter.HandleFunc("/delete/{id}", controller.DeleteAccount).Methods("DELETE")
	accountRouter.HandleFunc("/deposit/{id}", controller.DepositMoney).Methods("POST")
	accountRouter.HandleFunc("/withdraw/{id}", controller.WithdrawMoney).Methods("POST")
	accountRouter.HandleFunc("/transfer", controller.TransferMoney).Methods("POST")

	// Start the server on localhost:3000
	log.Printf("Server Live on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
