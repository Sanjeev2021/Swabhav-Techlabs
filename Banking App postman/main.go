package main

//controller "bankingApp/Controller"

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	controller "bankingApp/Controller"

)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	Handlefunction()
}

func Handlefunction() {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	router := mux.NewRouter()
	router = router.PathPrefix("/api/v1/bankingApp").Subrouter()

	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/create", controller.CreateUser).Methods("POST")
	userRouter.HandleFunc("/update", controller.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/delete", controller.DeleteUser).Methods("DELETE")
	userRouter.HandleFunc("/find", controller.FindUser).Methods("GET")
	userRouter.HandleFunc("/findall", controller.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/passbook", controller.GetUserPassbook).Methods("GET")

	bankRouter := router.PathPrefix("/bank").Subrouter()
	bankRouter.HandleFunc("/create", controller.CreateBank).Methods("POST")
	bankRouter.HandleFunc("/update", controller.UpdateBank).Methods("PUT")
	bankRouter.HandleFunc("/delete", controller.DeleteBank).Methods("DELETE")
	bankRouter.HandleFunc("/find", controller.GetBank).Methods("GET")
	bankRouter.HandleFunc("/findall", controller.GetAllBanks).Methods("GET")
	bankRouter.HandleFunc("/passbook", controller.GetBankPassbook).Methods("GET")

	accountRoute := router.PathPrefix("/account").Subrouter()
	accountRoute.HandleFunc("/create/{userid}", controller.CreateAccount).Methods("POST")
	accountRoute.HandleFunc("/update", controller.UpdateAccount).Methods("PUT")
	accountRoute.HandleFunc("/delete", controller.DeleteAccount).Methods("DELETE")
	accountRoute.HandleFunc("/find", controller.GetAccount).Methods("GET")
	accountRoute.HandleFunc("/findall", controller.GetAllAccounts).Methods("GET")
	accountRoute.HandleFunc("/passbook", controller.GetAccountPassbook).Methods("GET")

	accountRoute.HandleFunc("/withdraw", controller.WithdrawMoney).Methods("POST")
	accountRoute.HandleFunc("/deposit", controller.DepositMoney).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
