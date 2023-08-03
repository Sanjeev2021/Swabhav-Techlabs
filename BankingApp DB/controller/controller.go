package controller

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"bankingApp/Admin"
	"bankingApp/account"
	"bankingApp/bank"
	"bankingApp/services"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser *services.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser, err = services.CreateUser(newUser.FirstName, newUser.LastName, newUser.Username, newUser.Password)
	if err != nil {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&newUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	newUser, err := services.DeleteUser(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&newUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	var newUser *services.User

	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}
	newUser, err = services.UpdateUser(uint(id), newUser)
	if err != nil {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&newUser)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	newUser, err := services.GetUserById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newUser)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	users, err := services.GetAllUsers(page, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&users)
}

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var newadmin *Admin.Admin

	err := json.NewDecoder(r.Body).Decode(&newadmin)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}
	newadmin, err = Admin.CreateAdmin(newadmin.FirstName, newadmin.LastName, newadmin.Username, newadmin.Password)
	if err != nil {
		http.Error(w, "Admin already exist", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newadmin)
}

func GetAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid admin id", http.StatusBadRequest)
		return
	}

	newadmin, err := Admin.GetAdminById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newadmin)
}

func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid admin id", http.StatusBadRequest)
		return
	}

	newadmin, err := Admin.DeleteAdmin(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newadmin)
}

func UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid admin id", http.StatusBadRequest)
		return
	}

	var newadmin *Admin.Admin

	err = json.NewDecoder(r.Body).Decode(&newadmin)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	newadmin, err = Admin.UpdateAdmin(uint(id), newadmin)
	if err != nil {
		http.Error(w, "Could not update admin", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newadmin)
}

func GetAllAdmins(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	admins, err := Admin.GetAllAdmins(page, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&admins)
}

func CreateBank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminid, err := strconv.Atoi(vars["adminid"])
	if err != nil {
		http.Error(w, "Invalid admin id", http.StatusBadRequest)
		return
	}

	var newbank *bank.Bank

	err = json.NewDecoder(r.Body).Decode(&newbank)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	admin, err := Admin.GetAdminById(uint(adminid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newbank, err = bank.CreateBank(newbank.BankName, admin.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&newbank)
}

func GetBank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminid, err := strconv.Atoi(vars["adminid"])
	if err != nil {
		http.Error(w, "Invalid admin id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid bank id", http.StatusBadRequest)
		return
	}

	_, err = Admin.GetAdminById(uint(adminid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newbank, err := bank.GetBankById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&newbank)
}

func DeleteBank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminid, err := strconv.Atoi(vars["adminid"])
	if err != nil {
		http.Error(w, "Invalid admin id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid bank id", http.StatusBadRequest)
		return
	}

	admin, err := Admin.GetAdminById(uint(adminid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newbank, err := bank.GetBankById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if newbank.AdminID != admin.ID {
		http.Error(w, "Bank does not belong to admin", http.StatusBadRequest)
		return
	}

	newbank, err = bank.DeleteBank(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newbank)
}

func UpdateBank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adminid, err := strconv.Atoi(vars["adminid"])
	if err != nil {
		http.Error(w, "Invalid admin id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid bank id", http.StatusBadRequest)
		return
	}

	_, err = Admin.GetAdminById(uint(adminid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newbank *bank.Bank

	err = json.NewDecoder(r.Body).Decode(&newbank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newbank, err = bank.UpdateBank(uint(id), newbank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newbank)
}

func GetAllBanks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}
	adminid, err := strconv.Atoi(vars["adminid"])
	if err != nil {
		http.Error(w, "Invalid admin id", http.StatusBadRequest)
		return
	}

	_, err = Admin.GetAdminById(uint(adminid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	banks, err := bank.GetAllBanks(page, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&banks)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	var newaccount *account.Account

	err = json.NewDecoder(r.Body).Decode(&newaccount)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	_, err = services.GetUserById(uint(userid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = bank.GetBankById(newaccount.BankID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newaccount, err = account.CreateAccount(newaccount.AccountType, newaccount.Balance, newaccount.BankID, uint(userid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&newaccount)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {

		http.Error(w, "Invalid account id", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserById(uint(userid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newaccount, err := account.GetAccountById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if newaccount.UserID != user.ID {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&newaccount)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid account id", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserById(uint(userid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newaccount, err := account.GetAccountById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if newaccount.UserID != user.ID {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	newaccount, err = account.DeleteAccount(newaccount.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newaccount)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid account id", http.StatusBadRequest)
		return
	}

	_, err = services.GetUserById(uint(userid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newaccount *account.Account

	err = json.NewDecoder(r.Body).Decode(&newaccount)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	newaccount, err = account.UpdateAccount(uint(id), newaccount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newaccount)
}

func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	_, err = services.GetUserById(uint(userid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accounts, err := account.GetAllAccounts(page, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&accounts)
}

func TransferMoney(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	var newtransfer *account.AccountTranfer

	err = json.NewDecoder(r.Body).Decode(&newtransfer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.GetUserById(uint(userid)) //  to check weather the user transacting is his account or not
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newAccount, err := account.GetAccountById(newtransfer.FromAccountID) // to get the account from account id where we need to make transaction
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.ID != newAccount.UserID {
		http.Error(w, "Invalid user id", http.StatusBadRequest) // to check if user id and account id is same ..
		return
	}

	err = account.TransferMoney(newtransfer.FromAccountID, newtransfer.ToAccountID, newtransfer.Amount) // here we are calling transfer function
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&newtransfer)
}

func WithdrawMoney(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest) // to convert string to uint64
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid account id", http.StatusBadRequest)
		return
	}

	var newtransfer *account.AccountTranfer

	err = json.NewDecoder(r.Body).Decode(&newtransfer)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserById(uint(userid)) // we are getting the user from userid
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newAccount, err := account.GetAccountById(uint(id)) // we are getting account from accountid ,
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.ID != newAccount.UserID {
		http.Error(w, "Invalid user id", http.StatusBadRequest) // to check if user id and account id is same , so that no random person withdraws
		return
	}

	err = account.WithdrawMoney(uint(id), newtransfer.Amount) // calling function
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&newtransfer)
}

func DepositMoney(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid account id", http.StatusBadRequest)
		return
	}

	var newtransfer *account.AccountTranfer

	err = json.NewDecoder(r.Body).Decode(&newtransfer)
	if err != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserById(uint(userid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newAccount, err := account.GetAccountById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.ID != newAccount.UserID {
		http.Error(w, "Invalid user id", http.StatusBadRequest) // same above check here as well
		return
	}

	err = account.DepositMoney(uint(id), newtransfer.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // calling deposit function
		return
	}
	json.NewEncoder(w).Encode(&newtransfer)

}
