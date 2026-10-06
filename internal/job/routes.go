package job

import (
	"fmt"
	"net/http"

	"github.com/adrian-kurek/Job-Processing-Platform/common/request"
)

type jobHandler interface {
	Insert(w http.ResponseWriter, r *http.Request) error
}

const defaultPath = "/jobs"

type Route struct {
	jobHandler jobHandler
}

func NewRoute(jobHandler jobHandler) *Route {
	return &Route{
		jobHandler: jobHandler,
	}
}

func (rj Route) SetupRoutes(router *http.ServeMux) {
	router.Handle(fmt.Sprintf("POST %s/", defaultPath), request.Make(rj.jobHandler.Insert))
}
