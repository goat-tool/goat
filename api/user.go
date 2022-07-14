package api

import (
	// "encoding/json"
	// ut "github.com/go-playground/universal-translator"

	"encoding/json"
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
		//Todo logger log with prefix
	}
}

// func NewTestEndpoint(logger log.Logger, translator *ut.UniversalTranslator, validate *validator.Validate, service *test.Service) *TestEndpoint {
// 	return &TestEndpoint{
// 		logger:     logger.WithPrefix("api.test"),
// 		translator: translator,
// 		validate:   validate,
// 		service:    service,
// 	}
// }

func (e *UserEndpoint) Create(w http.ResponseWriter, r *http.Request) {
	//e.service.Log.Warn().Msg("TODO: CreateTest() in api/test.go")
	var input user.UserInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		respond(w, http.StatusBadRequest, "invalid body", nil)
		return
	}
	//fmt.Println("input", input)
	// err = e.validate.Struct(input)
	// if err != nil {
	// 	errs := getValidationError(err.(validator.ValidationErrors), trans)
	// 	respond(w, e.logger, http.StatusBadRequest, "validation failed", errs)
	// 	return
	// }

	createdUser, err := e.service.Create(input)
	if err != nil {
		// TODO add log + respond
		fmt.Println("error createdUser")
	}

	respond(w, http.StatusCreated, "test created successfully", createdUser)
}

func (e *UserEndpoint) GetAll(w http.ResponseWriter, _ *http.Request) {
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

func (e *UserEndpoint) GetById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	//fmt.Println("TODO: services/test/test.go GetTestById()")

	foundUser, err := e.service.Store.GetByID(id)
	if err != nil {
		switch err {
		case user.ErrIdParseFailed:
			respond(w, http.StatusBadRequest, err.Error(), nil)
		case user.ErrNotFound:
			respond(w, http.StatusNotFound, err.Error(), nil)
		default:
			respond(w, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	respond(w, http.StatusOK, "successfully found user", foundUser)
}

func (e *UserEndpoint) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var input *user.UserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		respond(w, http.StatusBadRequest, "invalid body", nil)
		return
	}

	// err = e.validate.Struct(input)
	// if err != nil {
	// 	errs := getValidationError(err.(validator.ValidationErrors), trans)
	// 	respond(w, e.logger, http.StatusBadRequest, "validation failed", errs)
	// 	return
	// }

	createdUser, err := e.service.Update(id, input)
	if err != nil {
		switch err {
		case user.ErrIdParseFailed:
			respond(w, http.StatusBadRequest, err.Error(), nil)
		case user.ErrNotFound:
			respond(w, http.StatusNotFound, err.Error(), nil)
		// case user.ErrPasswordChangeNotAllowed:
		// 	respond(w, http.StatusBadRequest, err.Error(), nil)
		default:
			respond(w, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	respond(w, http.StatusCreated, "user updated successfully", createdUser)
}

func (e *UserEndpoint) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := e.service.Delete(id)
	if err != nil {
		switch err {
		case user.ErrIdParseFailed:
			respond(w, http.StatusBadRequest, err.Error(), nil)
		case user.ErrNotFound:
			respond(w, http.StatusNotFound, err.Error(), nil)
		default:
			respond(w, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	respond(w, http.StatusOK, "successfully deleted", nil)
}
