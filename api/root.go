package api

import (
	"goat/services/root"

	"net/http"

	"git.bitcubix.io/go/log"
)

type RootEndpoint struct {
	service *root.Service
	log     *log.Logger
}

func NewRootEndpoint(log *log.Logger, service *root.Service) *RootEndpoint {
	return &RootEndpoint{
		service: service,
		log:     log,
	}
}

func (e *RootEndpoint) GetRoot(w http.ResponseWriter, _ *http.Request) {

	rootResponse := e.service.GetRoot()

	respond(w, e.log, http.StatusOK, "Welcome to goat api", rootResponse)
}
