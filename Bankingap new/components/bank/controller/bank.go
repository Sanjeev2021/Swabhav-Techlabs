package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"bankingapp/components/bank/service"
	"bankingapp/components/log"
	"bankingapp/errors"
	"bankingapp/models/bank"
	"bankingapp/web"
)

// BankController handles bank related requests
type BankController struct {
	service *service.BankService
	log     log.Log
}

// NewBankController returns new instance of BankController
func NewBankController(bankService *service.BankService, log log.Log) *BankController {
	return &BankController{
		service: bankService,
		log:     log,
	}
}

// RegisterRoutes registers routes for bank controller
func (controller *BankController) RegisterRoutes(router *mux.Router) {
	bankRouter := router.PathPrefix("/bank").Subrouter()
	bankRouter.HandleFunc("/create", controller.RegisterBank).Methods(http.MethodPost)
	bankRouter.HandleFunc("/", controller.GetAllBanks).Methods(http.MethodGet)
	bankRouter.HandleFunc("/{id}", controller.UpdateBank).Methods(http.MethodPut)
	bankRouter.HandleFunc("/{id}", controller.DeleteBank).Methods(http.MethodDelete)
}

// CreateBank creates a new bank
func (controller *BankController) RegisterBank(w http.ResponseWriter, r *http.Request) {
	newBank := bank.Bank{}
	// Unmarshal json.
	err := web.UnmarshalJSON(r, &newBank)
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	// Call add test method.
	err = controller.service.CreateBank(&newBank)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	// Writing Response with OK Status to ResposeWriter.
	web.RespondJSON(w, http.StatusCreated, newBank)
}

// GetAllBanks returns all banks
func (controller *BankController) GetAllBanks(w http.ResponseWriter, r *http.Request) {
	allBanks := &[]bank.Bank{}
	var totalCount int
	err := controller.service.GetAllBanks(allBanks, &totalCount)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSONWithXTotalCount(w, http.StatusOK, totalCount, allBanks)
}

// UpdateBank updates a bank
func (controller *BankController) UpdateBank(w http.ResponseWriter, r *http.Request) {
	bankToUpdate := bank.Bank{}
	err := web.UnmarshalJSON(r, &bankToUpdate)
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
	bankToUpdate.ID = uint(id)
	err = controller.service.UpdateBank(&bankToUpdate)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, bankToUpdate)
}

// DeleteBank deletes a bank
func (controller *BankController) DeleteBank(w http.ResponseWriter, r *http.Request) {

	bankToDelete := bank.Bank{}
	var err error
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	bankToDelete.ID = uint(id)
	err = controller.service.DeleteBank(&bankToDelete)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return

	}
	web.RespondJSON(w, http.StatusOK, "Bank Deleted Successfully")

}
