package api

// "encoding/json"
// ut "github.com/go-playground/universal-translator"
// "net/http"

// "goat/log"
import "goat/services/user"

// "github.com/go-playground/validator/v10"
// "github.com/gorilla/mux"

type UserEndpoint struct {
	// logger     log.Logger
	// translator *ut.UniversalTranslator
	// validate   *validator.Validate
	service *user.Service
}

func NewUserEndpoint(service *user.Service) *UserEndpoint {
	return &UserEndpoint{
		service: service,
	}
}

// func NewUserEndpoint(logger log.Logger, translator *ut.UniversalTranslator, validate *validator.Validate, service *user.Service) *UserEndpoint {
// 	return &UserEndpoint{
// 		logger:     logger.WithPrefix("api.user"),
// 		translator: translator,
// 		validate:   validate,
// 		service:    service,
// 	}
// }

// func (e *UserEndpoint) CreateUser(w http.ResponseWriter, r *http.Request) {
// 	trans := getTranslator(r, e.translator)

// 	var input user.UserInput

// 	err := json.NewDecoder(r.Body).Decode(&input)
// 	if err != nil {
// 		respond(w, e.logger, http.StatusBadRequest, "invalid body", nil)
// 		return
// 	}

// 	err = e.validate.Struct(input)
// 	if err != nil {
// 		errs := getValidationError(err.(validator.ValidationErrors), trans)
// 		respond(w, e.logger, http.StatusBadRequest, "validation failed", errs)
// 		return
// 	}

// 	createdUser, err := e.service.Create(input)
// 	if err != nil {
// 		// TODO add error handling
// 	}

// 	respond(w, e.logger, http.StatusCreated, "user created successfully", createdUser)
// }

// func (e *UserEndpoint) GetAllUsers(w http.ResponseWriter, _ *http.Request) {
// 	users, err := e.service.GetAll()
// 	if err != nil {
// 		switch err {
// 		case user.ErrNotFound:
// 			respond(w, e.logger, http.StatusNotFound, err.Error(), nil)
// 		default:
// 			respond(w, e.logger, http.StatusInternalServerError, err.Error(), nil)
// 		}
// 		return
// 	}

// 	respond(w, e.logger, http.StatusOK, "load all users successfully", users)
// }

// func (e *UserEndpoint) GetUserById(w http.ResponseWriter, r *http.Request) {
// 	id := mux.Vars(r)["id"]

// 	foundUser, err := e.service.Store.GetByID(id)
// 	if err != nil {
// 		switch err {
// 		case user.ErrIdParseFailed:
// 			respond(w, e.logger, http.StatusBadRequest, err.Error(), nil)
// 		case user.ErrNotFound:
// 			respond(w, e.logger, http.StatusNotFound, err.Error(), nil)
// 		default:
// 			respond(w, e.logger, http.StatusInternalServerError, err.Error(), nil)
// 		}
// 		return
// 	}

// 	respond(w, e.logger, http.StatusOK, "successfully found user", foundUser)
// }

// func (e *UserEndpoint) UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	var input user.UserInput
// 	id := mux.Vars(r)["id"]

// 	err := json.NewDecoder(r.Body).Decode(&input)
// 	if err != nil {
// 		respond(w, e.logger, http.StatusBadRequest, "invalid body", nil)
// 		return
// 	}

// 	updatedUser, err := e.service.Update(id, input)
// 	if err != nil {
// 		switch err {
// 		case user.ErrIdParseFailed:
// 			respond(w, e.logger, http.StatusBadRequest, err.Error(), nil)
// 		case user.ErrNotFound:
// 			respond(w, e.logger, http.StatusNotFound, err.Error(), nil)
// 		case user.ErrPasswordChangeNotAllowed:
// 			respond(w, e.logger, http.StatusBadRequest, err.Error(), nil)
// 		default:
// 			respond(w, e.logger, http.StatusInternalServerError, err.Error(), nil)
// 		}
// 		return
// 	}

// 	respond(w, e.logger, http.StatusOK, "successfully updated user", updatedUser)
// }

// func (e *UserEndpoint) DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	id := mux.Vars(r)["id"]

// 	err := e.service.Delete(id)
// 	if err != nil {
// 		switch err {
// 		case user.ErrIdParseFailed:
// 			respond(w, e.logger, http.StatusBadRequest, err.Error(), nil)
// 		case user.ErrNotFound:
// 			respond(w, e.logger, http.StatusNotFound, err.Error(), nil)
// 		default:
// 			respond(w, e.logger, http.StatusInternalServerError, err.Error(), nil)
// 		}
// 		return
// 	}

// 	respond(w, e.logger, http.StatusOK, "successfully deleted", nil)
// }
