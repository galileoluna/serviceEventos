package app

import (
	"net/http"

	"github.com/galileoluna/serviceEventos/controllers"
)

func mapUrls() {

	eventosRouter := router.PathPrefix("/evento").Subrouter()
	eventosRouter.Path("").Methods(http.MethodPost).HandlerFunc(controllers.EventoController.InsertEvento)
	eventosRouter.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(controllers.EventoController.DeleteEvento)
	eventosRouter.Path("/{id}").Methods(http.MethodPut).HandlerFunc(controllers.EventoController.UpdateEvento)
	eventosRouter.Path("/data/{id}").Methods(http.MethodGet).HandlerFunc(controllers.EventoController.GetEvento)
	eventosRouter.Path("/all").Methods(http.MethodGet).HandlerFunc(controllers.EventoController.GetEventos)
}
