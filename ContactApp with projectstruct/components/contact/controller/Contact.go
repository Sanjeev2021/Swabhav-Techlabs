package controller

import (
	"contactApp/components/contact/service"
	"contactApp/components/log"
	"contactApp/web"
	"net/http"

	"github.com/gorilla/mux"
)

// ContactController gives access to CRUD operations for entity
type ContactController struct {
	log     log.Log
	service *service.ContactService
}

// NewContactController returns new instance of ContactController
func NewContactController(contactService *service.ContactService,
	log log.Log) *ContactController {
	return &ContactController{
		service: contactService,
		log:     log,
	}
}
func (controller *ContactController) RegisterRoutes(router *mux.Router) {
	// userRouter := router.PathPrefix("/user").Subrouter()
	// userRouter.HandleFunc("/register", controller.CreateContact).Methods(http.MethodPost)
	// userRouter.HandleFunc("/", controller.GetContact).Methods(http.MethodGet)
	// userRouter.HandleFunc("/:id", controller.UpdateContact).Methods(http.MethodPut)
	// userRouter.HandleFunc("/:id", controller.DeleteContact).Methods(http.MethodDelete)

}
func (controller *ContactController) CreateContact(w http.ResponseWriter, r *http.Request) {
	web.RespondJSON(w, http.StatusOK, "CreateContact successfull.")
}
func (controller *ContactController) GetContact(w http.ResponseWriter, r *http.Request) {
	web.RespondJSON(w, http.StatusOK, "GetContact successfull.")
}
func (controller *ContactController) UpdateContact(w http.ResponseWriter, r *http.Request) {
	web.RespondJSON(w, http.StatusOK, "UpdateContact successfull.")
}
func (controller *ContactController) DeleteContact(w http.ResponseWriter, r *http.Request) {
	web.RespondJSON(w, http.StatusOK, "DeleteContact successfull.")
}
