package controller

import (
	"contactApp/components/contact_info/service"
	"contactApp/components/log"
	"contactApp/web"
	"net/http"

	"github.com/gorilla/mux"
)

// ContactInfoController gives access to CRUD operations for entity
type ContactInfoController struct {
	log     log.Log
	service *service.ContactInfoService
}

// NewContactInfoController returns new instance of ContactInfoController
func NewContactInfoController(contactService *service.ContactInfoService,
	log log.Log) *ContactInfoController {
	return &ContactInfoController{
		service: contactService,
		log:     log,
	}
}
func (controller *ContactInfoController) RegisterRoutes(router *mux.Router) {
	// userRouter := router.PathPrefix("/user").Subrouter()
	// userRouter.HandleFunc("/register", controller.CreateContactInfo).Methods(http.MethodPost)
	// userRouter.HandleFunc("/", controller.GetContactInfo).Methods(http.MethodGet)
	// userRouter.HandleFunc("/:id", controller.UpdateContactInfo).Methods(http.MethodPut)
	// userRouter.HandleFunc("/:id", controller.DeleteContactInfo).Methods(http.MethodDelete)

}
func (controller *ContactInfoController) CreateContactInfo(w http.ResponseWriter, r *http.Request) {
	web.RespondJSON(w, http.StatusOK, "CreateContactInfo successfull.")
}
func (controller *ContactInfoController) GetContactInfo(w http.ResponseWriter, r *http.Request) {
	web.RespondJSON(w, http.StatusOK, "GetContactInfo successfull.")
}
func (controller *ContactInfoController) UpdateContactInfo(w http.ResponseWriter, r *http.Request) {
	web.RespondJSON(w, http.StatusOK, "UpdateContactInfo successfull.")
}
func (controller *ContactInfoController) DeleteContactInfo(w http.ResponseWriter, r *http.Request) {
	web.RespondJSON(w, http.StatusOK, "DeleteContactInfo successfull.")
}