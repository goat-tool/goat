package api

import (
	// "encoding/json"
	// ut "github.com/go-playground/universal-translator"

	"encoding/json"
	"net/http"

	"goat/log"
	"goat/services/test"

	"github.com/gorilla/mux"
	// "github.com/go-playground/validator/v10"
	// "github.com/gorilla/mux"
)

type TestEndpoint struct {
	log *log.Logger
	// translator *ut.UniversalTranslator
	// validate   *validator.Validate
	service *test.Service
}

func NewTestEndpoint(log *log.Logger, service *test.Service) *TestEndpoint {
	return &TestEndpoint{
		service: service,
		log:     log,
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

func (e *TestEndpoint) Create(w http.ResponseWriter, r *http.Request) {
	var input test.TestInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		respond(w, e.log, http.StatusBadRequest, "invalid body", nil)
		return
	}
	//fmt.Println("input", input)
	// err = e.validate.Struct(input)
	// if err != nil {
	// 	errs := getValidationError(err.(validator.ValidationErrors), trans)
	// 	respond(w, e.logger, http.StatusBadRequest, "validation failed", errs)
	// 	return
	// }

	createdTest, err := e.service.Create(input)
	if err != nil {
		respond(w, e.log, http.StatusBadRequest, "invalid body", nil)
	}

	respond(w, e.log, http.StatusCreated, "test created successfully", createdTest)
}

func (e *TestEndpoint) GetAll(w http.ResponseWriter, _ *http.Request) {
	tests, err := e.service.GetAll()
	if err != nil {
		switch err {
		case test.ErrNotFound:
			respond(w, e.log, http.StatusNotFound, err.Error(), nil)
		default:
			respond(w, e.log, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	respond(w, e.log, http.StatusOK, "load all tests successfully", tests)
}

func (e *TestEndpoint) GetById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	foundTest, err := e.service.Store.GetByID(id)
	if err != nil {
		switch err {
		case test.ErrIdParseFailed:
			respond(w, e.log, http.StatusBadRequest, err.Error(), nil)
		case test.ErrNotFound:
			respond(w, e.log, http.StatusNotFound, err.Error(), nil)
		default:
			respond(w, e.log, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	respond(w, e.log, http.StatusOK, "successfully found test", foundTest)
}

func (e *TestEndpoint) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var input *test.TestInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		respond(w, e.log, http.StatusBadRequest, "invalid body", nil)
		return
	}

	// err = e.validate.Struct(input)
	// if err != nil {
	// 	errs := getValidationError(err.(validator.ValidationErrors), trans)
	// 	respond(w, e.logger, http.StatusBadRequest, "validation failed", errs)
	// 	return
	// }

	createdTest, err := e.service.Update(id, input)
	if err != nil {
		switch err {
		case test.ErrIdParseFailed:
			respond(w, e.log, http.StatusBadRequest, err.Error(), nil)
		case test.ErrNotFound:
			respond(w, e.log, http.StatusNotFound, err.Error(), nil)
		default:
			respond(w, e.log, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	respond(w, e.log, http.StatusCreated, "test updated successfully", createdTest)
}

func (e *TestEndpoint) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := e.service.Delete(id)
	if err != nil {
		switch err {
		case test.ErrIdParseFailed:
			respond(w, e.log, http.StatusBadRequest, err.Error(), nil)
		case test.ErrNotFound:
			respond(w, e.log, http.StatusNotFound, err.Error(), nil)
		default:
			respond(w, e.log, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	respond(w, e.log, http.StatusOK, "successfully deleted", nil)
}
