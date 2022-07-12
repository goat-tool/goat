package api

import (
	"encoding/json"
	"net/http"

	// "goat/log"

	// ut "github.com/go-playground/universal-translator"
	// "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Api struct {
	Health *HealthEndpoint
	User   *UserEndpoint
	Test   *TestEndpoint
}

type Body struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// type ValidationError map[string]string
// type ValidationErrors []ValidationError

func (a *Api) Setup(router *mux.Router) {
	v1 := router.PathPrefix("/v1").Subrouter()

	// Health
	v1.HandleFunc("/health", a.Health.GetHealth).Methods(http.MethodGet)

	// User
	// v1.HandleFunc("/users", a.User.CreateUser).Methods(http.MethodPost)
	v1.HandleFunc("/users", a.User.GetAllUsers).Methods(http.MethodGet)
	v1.HandleFunc("/users/{id}", a.User.GetUserById).Methods(http.MethodGet)
	// v1.HandleFunc("/users/{id}", a.User.UpdateUser).Methods(http.MethodPut)
	// v1.HandleFunc("/users/{id}", a.User.DeleteUser).Methods(http.MethodDelete)

	// Test
	v1.HandleFunc("/test", a.Test.CreateTest).Methods(http.MethodPost)
	v1.HandleFunc("/tests", a.Test.GetAllTests).Methods(http.MethodGet)
	v1.HandleFunc("/tests/{id}", a.Test.GetTestById).Methods(http.MethodGet)
	v1.HandleFunc("/tests/{id}", a.Test.UpdateTest).Methods(http.MethodPut)
	v1.HandleFunc("/tests/{id}", a.Test.DeleteTest).Methods(http.MethodDelete)

}

func respond(w http.ResponseWriter, status int, message string, data interface{}) {
	body := Body{}
	body.Status = status

	if message != "" {
		body.Message = message
	}

	if data != nil {
		body.Data = data
	}

	bodyString, err := json.Marshal(body)
	if err != nil {
		// logger.Errorf("fail to parse response body json: %v", err)
		respond(w, http.StatusInternalServerError, "internal error", nil)
		return
	}

	w.WriteHeader(status)
	_, err = w.Write(bodyString)
	if err != nil {
		// logger.Errorf("fail to write response body: %v", err)
		return
	}
}

// func getTranslator(r *http.Request, translator *ut.UniversalTranslator) ut.Translator {
// 	lang := r.Header.Get("Accept-Language")
// 	trans, _ := translator.GetTranslator(lang)
// 	return trans
// }

// func getValidationError(err validator.ValidationErrors, translator ut.Translator) ValidationErrors {
// 	ve := ValidationErrors{}

// 	for _, e := range err {
// 		ve = append(ve, ValidationError{e.Field(): e.Translate(translator)})
// 	}

// 	return ve
// }

// func respond(w http.ResponseWriter, logger log.Logger, status int, message string, data interface{}) {
// 	body := Body{}
// 	body.Status = status

// 	if message != "" {
// 		body.Message = message
// 	}

// 	if data != nil {
// 		body.Data = data
// 	}

// 	bodyString, err := json.Marshal(body)
// 	if err != nil {
// 		logger.Errorf("fail to parse response body json: %v", err)
// 		respond(w, logger, http.StatusInternalServerError, "internal error", nil)
// 		return
// 	}

// 	w.WriteHeader(status)
// 	_, err = w.Write(bodyString)
// 	if err != nil {
// 		logger.Errorf("fail to write response body: %v", err)
// 		return
// 	}
// }

// func getTranslator(r *http.Request, translator *ut.UniversalTranslator) ut.Translator {
// 	lang := r.Header.Get("Accept-Language")
// 	trans, _ := translator.GetTranslator(lang)
// 	return trans
// }

// func getValidationError(err validator.ValidationErrors, translator ut.Translator) ValidationErrors {
// 	ve := ValidationErrors{}

// 	for _, e := range err {
// 		ve = append(ve, ValidationError{e.Field(): e.Translate(translator)})
// 	}

// 	return ve
// }
