package api

import (
	// "encoding/json"
	// ut "github.com/go-playground/universal-translator"

	"fmt"
	"net/http"

	// "goat/log"
	"goat/services/user"

	"github.com/gorilla/mux"
	// "github.com/go-playground/validator/v10"
	// "github.com/gorilla/mux"
)

type UserEndpoint struct {
	// logger     log.Logger
	// translator *ut.UniversalTranslator
	// validate   *validator.Validate
	service *user.Service
}

func NewUserEndpoint(service *user.Service) *UserEndpoint {
	return &UserEndpoint{
		service: service,
		//logger: log with prefix
	}
}

func (e *UserEndpoint) GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := e.service.GetAll()
	if err != nil {
		switch err {
		case user.ErrNotFound:
			respond(w, http.StatusNotFound, err.Error(), nil)
		default:
			respond(w, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	respond(w, http.StatusOK, "load all users successfully", users)
}

func (e *UserEndpoint) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	fmt.Println("TODO: services/user/user.go GetUserById()")

	foundUser := id

	respond(w, http.StatusOK, "successfully found user", foundUser)
}
