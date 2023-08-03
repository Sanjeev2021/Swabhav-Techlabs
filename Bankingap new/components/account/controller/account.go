package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"bankingapp/components/account/service" // Import the service package here
	"bankingapp/components/log"
	"bankingapp/errors"
	"bankingapp/models/account"
	"bankingapp/web"
)

// AccountController is a struct that defines the dependencies of the AccountController
type AccountController struct {
	service *service.AccountService
	log     log.Log
}

// NewAccountController returns new instance of AccountController
func NewAccountController(accountService *service.AccountService, log log.Log) *AccountController {
	return &AccountController{
		service: accountService,
		log:     log,
	}
}

// RegisterRoutes registers routes for account controller
func (controller *AccountController) RegisterRoutes(router *mux.Router) {
	accountRouter := router.PathPrefix("/account").Subrouter()
	accountRouter.HandleFunc("/register", controller.RegisterAccount).Methods(http.MethodPost)
	accountRouter.HandleFunc("/", controller.GetAllAccounts).Methods(http.MethodGet)
	accountRouter.HandleFunc("/{id}", controller.UpdateAccount).Methods(http.MethodPut)
	accountRouter.HandleFunc("/{id}", controller.DeleteAccount).Methods(http.MethodDelete)
}

// register account
func (controller *AccountController) RegisterAccount(w http.ResponseWriter, r *http.Request) {
	newAccount := account.Account{}
	err := web.UnmarshalJSON(r, &newAccount)
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	err = controller.service.CreateAccount(&newAccount)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusCreated, newAccount)
}

// get all accounts
func (controller *AccountController) GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	allAccounts := &[]account.Account{}
	var totalCount int
	err := controller.service.GetAllAccounts(allAccounts, &totalCount)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, allAccounts)
}

// Update Account
func (controller *AccountController) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	updatedAccount := account.Account{}
	err := web.UnmarshalJSON(r, &updatedAccount)
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	updatedAccount.ID = uint(id)
	err = controller.service.UpdateAccount(&updatedAccount)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, updatedAccount)
}

// Delete Account
func (controller *AccountController) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	accountToDelete := account.Account{}
	var err error
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	accountToDelete.ID = uint(id)
	err = controller.service.DeleteAccount(&accountToDelete)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, accountToDelete)
}
