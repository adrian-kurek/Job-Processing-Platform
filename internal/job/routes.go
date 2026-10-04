package job

import (
	"net/http"
)

type jobHandler interface {
	Insert(w http.ResponseWriter, r *http.Request) error
}

type Route struct {
	jobHandler jobHandler
}

func NewRoute(jobHandler jobHandler) *Route {
	return &Route{
		jobHandler: jobHandler,
	}
}
