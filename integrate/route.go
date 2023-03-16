package integrate

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/jtfm/ec2-scheduler/internal/render"
)

func Router() http.Handler {
	router := chi.NewRouter()
	router.HandleFunc("/", render.Instances)
	return router
}
