package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	//"bankingapp/app"
	"bankingapp/components/admin/service"
	"bankingapp/components/log"
	"bankingapp/errors"
	"bankingapp/models/admin"
	"bankingapp/web"
)

// AdminController gives access to CRUD operations for entity
type AdminController struct {
	log     log.Log
	service *service.AdminService
}

// NewAdminController returns new instance of AdminController
func NewAdminController(adminService *service.AdminService, log log.Log) *AdminController {
	return &AdminController{
		service: adminService,
		log:     log,
	}
}
func (controller *AdminController) RegisterRoutes(router *mux.Router) {
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/register", controller.registerAdmin).Methods(http.MethodPost)
	adminRouter.HandleFunc("/", controller.GetAllAdmins).Methods(http.MethodGet)
	adminRouter.HandleFunc("/{id}", controller.UpdateAdmin).Methods(http.MethodPut)
	adminRouter.HandleFunc("/{id}", controller.DeleteAdmin).Methods(http.MethodDelete)
	controller.log.Print("==============================adminRegisterRoutes==========================")
}
func (controller *AdminController) registerAdmin(w http.ResponseWriter, r *http.Request) {
	newAdmin := admin.Admin{}
	// Unmarshal json.
	err := web.UnmarshalJSON(r, &newAdmin)
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	// Call add test method.
	err = controller.service.CreateAdmin(&newAdmin)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	// Writing Response with OK Status to ResposeWriter.
	web.RespondJSON(w, http.StatusCreated, newAdmin)
}
func (controller *AdminController) GetAllAdmins(w http.ResponseWriter, r *http.Request) {
	allAdmins := &[]admin.Admin{}
	var totalCount int
	err := controller.service.GetAllAdmins(allAdmins, &totalCount)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	// Writing Response with OK Status to ResonseWriter,
	web.RespondJSONWithXTotalCount(w, http.StatusOK, totalCount, allAdmins)
}
func (controller *AdminController) UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("==============================adminToUpdate==========================")
	adminToUpdate := admin.Admin{}

	// Unmarshal JSON.
	fmt.Println(r.Body)
	err := web.UnmarshalJSON(r, &adminToUpdate)
	if err != nil {
		fmt.Println("==============================err from UnmarshalJSON==========================")
		controller.log.Print(err.Error())
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	vars := mux.Vars(r)

	intID, err := strconv.Atoi(vars["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	adminToUpdate.ID = uint(intID)
	fmt.Println("==============================adminToUpdate==========================")
	fmt.Println(&adminToUpdate)
	// Call update test method.
	err = controller.service.UpdateAdmin(&adminToUpdate)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}

	web.RespondJSON(w, http.StatusOK, adminToUpdate)
}

func (controller *AdminController) DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("==============================DeleteAdmin call==========================")
	adminToDelete := admin.Admin{}
	var err error
	vars := mux.Vars(r)
	intID, err := strconv.Atoi(vars["id"])
	if err != nil {
		controller.log.Print(err)
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	adminToDelete.ID = uint(intID)
	err = controller.service.DeleteAdmin(&adminToDelete)
	if err != nil {
		controller.log.Print(err.Error())
		web.RespondError(w, err)
		return
	}
	web.RespondJSON(w, http.StatusOK, "Delete admin successfull.")
}
