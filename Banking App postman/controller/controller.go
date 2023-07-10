package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	account "bankingApp/account"
	"bankingApp/bank"
	"bankingApp/user"
)

type UserUpdate struct {
	ID    string
	Field string
	Value string
}

type BankUpdate struct {
	ID    string
	Field string
	Value string
}

type AccountUpdate struct {
	ID    string
	Field string
	Value string
}

type WithdrawMoneyAccount struct {
	ID     string
	Amount float64
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userValues *user.User

	err := json.NewDecoder(r.Body).Decode(&userValues)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	newUser, err := user.CreateUser(userValues.FirstName, userValues.LastName, userValues.Username, userValues.IsAdmin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&newUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userValues *UserUpdate

	err := json.NewDecoder(r.Body).Decode(&userValues)
	if err != nil {
		http.Error(w, "Missing Parameters", http.StatusBadRequest)
		return
	}

	newUser, err := user.UpdateUser(userValues.ID, userValues.Field, userValues.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var userValues *UserUpdate

	err := json.NewDecoder(r.Body).Decode(&userValues)
	if err != nil {
		http.Error(w, "Missing Parameters", http.StatusBadRequest)
		return
	}

	deletedUser, err := user.DeleteUser(userValues.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&deletedUser)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	var userValues *UserUpdate

	err := json.NewDecoder(r.Body).Decode(&userValues)
	if err != nil {
		http.Error(w, "Missing Parameters", http.StatusBadRequest)
		return
	}

	foundUser, userExist := user.FindUser(user.UsersCreatedByMe, userValues.ID)
	if !userExist {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&foundUser)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&user.UsersCreatedByMe)
}

func CreateBank(w http.ResponseWriter, r *http.Request) {
	var bankValues *bank.Bank

	err := json.NewDecoder(r.Body).Decode(&bankValues)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	newBank := bank.CreateBank(bankValues.BankName)
	json.NewEncoder(w).Encode(&newBank)
}

func GetBank(w http.ResponseWriter, r *http.Request) {
	var bankValues *BankUpdate

	err := json.NewDecoder(r.Body).Decode(&bankValues)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	foundBank, bankExist := bank.FindBank(bank.BanksCreatedByMe, bankValues.ID)
	if !bankExist {
		http.Error(w, "Bank does not exist", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&foundBank)
}

func UpdateBank(w http.ResponseWriter, r *http.Request) {
	var bankValues *BankUpdate

	err := json.NewDecoder(r.Body).Decode(&bankValues)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	updatedBank, err := bank.UpdateBank(bankValues.ID, bankValues.Field, bankValues.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&updatedBank)
}

func DeleteBank(w http.ResponseWriter, r *http.Request) {
	var bankValues *BankUpdate

	err := json.NewDecoder(r.Body).Decode(&bankValues)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	deletedBank, err := bank.DeleteBank(bankValues.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&deletedBank)
}

func GetAllBanks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&bank.BanksCreatedByMe)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["userid"]

	var accountValues *account.Account

	err := json.NewDecoder(r.Body).Decode(&accountValues)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	bankFound, bankExists := bank.FindBankByName(bank.BanksCreatedByMe, accountValues.BankName)
	if !bankExists {
		http.Error(w, "Bank does not exist", http.StatusBadRequest)
		return
	}

	userFound, userExists := user.FindUser(user.UsersCreatedByMe, userid)

	if !userExists {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}

	newAccount, err := account.CreateAccount(accountValues.BankName, accountValues.AccountType, accountValues.AccountBalance, userFound)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bank.AddAccount(bankFound, newAccount)
	json.NewEncoder(w).Encode(&newAccount)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var accountValues *AccountUpdate

	err := json.NewDecoder(r.Body).Decode(&accountValues)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	updatedAccount, err := account.UpdateAccount(accountValues.ID, accountValues.Field, accountValues.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(&updatedAccount)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	var accountValues *AccountUpdate

	err := json.NewDecoder(r.Body).Decode(&accountValues)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	deletedAccount, err := account.DeleteAccount(accountValues.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&deletedAccount)
}

func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&account.AccountsCreatedByMe)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	var accountValues *AccountUpdate

	err := json.NewDecoder(r.Body).Decode(&accountValues)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	foundAccount, accountExist := account.FindAccount(account.AccountsCreatedByMe, accountValues.ID)
	if !accountExist {
		http.Error(w, "Account does not exist", http.StatusBadRequest)
		return

	}
	json.NewEncoder(w).Encode(&foundAccount)
}

func WithdrawMoney(w http.ResponseWriter, r *http.Request) {
	var accountValues *WithdrawMoneyAccount

	err := json.NewDecoder(r.Body).Decode(&accountValues)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	accountFound, _ := account.FindAccount(account.AccountsCreatedByMe, accountValues.ID)
	bank, _ := bank.FindBankByName(bank.BanksCreatedByMe, accountFound.BankName)

	err = accountFound.WithdrawMoney(accountValues.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	name := accountFound.Owner.FirstName + " " + accountFound.Owner.LastName
	amount := strconv.FormatFloat(accountValues.Amount, 'f', -1, 64)
	currentTime := account.GetCurrentTime()
	statement := currentTime + " " + name + " withdrew $" + amount + " from " + accountFound.BankName + " " + accountFound.AccountType + " account"

	accountFound.PassBook = append(accountFound.PassBook, statement)
	bank.PassBook = append(bank.PassBook, statement)
	accountFound.Owner.PassBook = append(accountFound.Owner.PassBook, statement)

	type Response struct {
		Confirmation string
	}

	json.NewEncoder(w).Encode(&Response{Confirmation: statement})
}

func DepositMoney(w http.ResponseWriter, r *http.Request) {
	var accountValues *WithdrawMoneyAccount

	err := json.NewDecoder(r.Body).Decode(&accountValues)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	accountFound, _ := account.FindAccount(account.AccountsCreatedByMe, accountValues.ID)
	bank, _ := bank.FindBankByName(bank.BanksCreatedByMe, accountFound.BankName)

	accountFound.DepositMoney(accountValues.Amount)

	name := accountFound.Owner.FirstName + " " + accountFound.Owner.LastName
	amount := strconv.FormatFloat(accountValues.Amount, 'f', -1, 64)
	currentTime := account.GetCurrentTime()
	statement := currentTime + " " + name + " deposited $" + amount + " to " + accountFound.BankName + " " + accountFound.AccountType + " account"

	accountFound.PassBook = append(accountFound.PassBook, statement)
	bank.PassBook = append(bank.PassBook, statement)
	accountFound.Owner.PassBook = append(accountFound.Owner.PassBook, statement)

	type Response struct {
		Confirmation string
	}

	json.NewEncoder(w).Encode(&Response{Confirmation: statement})
}

func GetUserPassbook(w http.ResponseWriter, r *http.Request) {
	var userValues *UserUpdate

	err := json.NewDecoder(r.Body).Decode(&userValues)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	userFound, userExists := user.FindUser(user.UsersCreatedByMe, userValues.ID)

	if !userExists {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&userFound.PassBook)
}

func GetAccountPassbook(w http.ResponseWriter, r *http.Request) {
	var accountValues *AccountUpdate

	err := json.NewDecoder(r.Body).Decode(&accountValues)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	accountFound, accountExists := account.FindAccount(account.AccountsCreatedByMe, accountValues.ID)

	if !accountExists {
		http.Error(w, "Account does not exist", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&accountFound.PassBook)
}

func GetBankPassbook(w http.ResponseWriter, r *http.Request) {
	var bankValues *bank.Bank

	err := json.NewDecoder(r.Body).Decode(&bankValues)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	bankFound, bankExists := bank.FindBankByName(bank.BanksCreatedByMe, bankValues.BankName)

	if !bankExists {
		http.Error(w, "Bank does not exist", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&bankFound.PassBook)
}
