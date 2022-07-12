package api

import (
	// "encoding/json"
	// ut "github.com/go-playground/universal-translator"

	"encoding/json"
	"fmt"
	"net/http"

	// "goat/log"
	"goat/services/test"

	"github.com/gorilla/mux"
	// "github.com/go-playground/validator/v10"
	// "github.com/gorilla/mux"
)

type TestEndpoint struct {
	// logger     log.Logger
	// translator *ut.UniversalTranslator
	// validate   *validator.Validate
	service *test.Service
}

func NewTestEndpoint(service *test.Service) *TestEndpoint {
	return &TestEndpoint{
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

func (e *TestEndpoint) CreateTest(w http.ResponseWriter, r *http.Request) {
	e.service.Log.Warn().Msg("TODO: CreateTest() in api/test.go")
	var input test.TestInput

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

	createdTest, err := e.service.Create(input)
	if err != nil {
		// TODO add log + respond
		fmt.Println("error createdTest")
	}

	respond(w, http.StatusCreated, "test created successfully", createdTest)
}

func (e *TestEndpoint) GetAllTests(w http.ResponseWriter, _ *http.Request) {
	tests, err := e.service.GetAll()
	if err != nil {
		switch err {
		case test.ErrNotFound:
			respond(w, http.StatusNotFound, err.Error(), nil)
		default:
			respond(w, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	respond(w, http.StatusOK, "load all tests successfully", tests)
}

func (e *TestEndpoint) GetTestById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	//fmt.Println("TODO: services/test/test.go GetTestById()")

	foundTest, err := e.service.Store.GetByID(id)
	if err != nil {
		switch err {
		case test.ErrIdParseFailed:
			respond(w, http.StatusBadRequest, err.Error(), nil)
		case test.ErrNotFound:
			respond(w, http.StatusNotFound, err.Error(), nil)
		default:
			respond(w, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	respond(w, http.StatusOK, "successfully found test", foundTest)
}

func (e *TestEndpoint) UpdateTest(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var input *test.TestInput
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

	createdTest, err := e.service.Update(id, input)
	if err != nil {
		switch err {
		case test.ErrIdParseFailed:
			respond(w, http.StatusBadRequest, err.Error(), nil)
		case test.ErrNotFound:
			respond(w, http.StatusNotFound, err.Error(), nil)
		// case test.ErrPasswordChangeNotAllowed:
		// 	respond(w, http.StatusBadRequest, err.Error(), nil)
		default:
			respond(w, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	respond(w, http.StatusCreated, "test updated successfully", createdTest)
}

func (e *TestEndpoint) DeleteTest(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := e.service.Delete(id)
	if err != nil {
		switch err {
		case test.ErrIdParseFailed:
			respond(w, http.StatusBadRequest, err.Error(), nil)
		case test.ErrNotFound:
			respond(w, http.StatusNotFound, err.Error(), nil)
		default:
			respond(w, http.StatusInternalServerError, err.Error(), nil)
		}
		return
	}

	respond(w, http.StatusOK, "successfully deleted", nil)
}
