package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/galileoluna/serviceEventos/domain/evento"
	"github.com/galileoluna/serviceEventos/services"
	"github.com/gorilla/mux"
)

var (
	EventoController eventoControllerInterface = &eventoController{}
)

type eventoControllerInterface interface {
	InsertEvento(w http.ResponseWriter, r *http.Request)
	GetEvento(w http.ResponseWriter, r *http.Request)
	UpdateEvento(w http.ResponseWriter, r *http.Request)
	GetEventos(w http.ResponseWriter, r *http.Request)
	DeleteEvento(w http.ResponseWriter, r *http.Request)
}

type eventoController struct {
}

func (c *eventoController) InsertEvento(w http.ResponseWriter, r *http.Request) {
	var nuevoEvento evento.Evento
	err := json.NewDecoder(r.Body).Decode(&nuevoEvento)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	services.EventoService.InsertEvento(nuevoEvento)
	w.WriteHeader(http.StatusCreated)
}
func (c *eventoController) GetEvento(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	ID_Evento, _ := strconv.ParseInt(id, 10, 64)
	eventoRequerido, _ := services.EventoService.GetEvento(ID_Evento)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(eventoRequerido); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func (c *eventoController) UpdateEvento(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	ID64, _ := strconv.ParseInt(id, 10, 64)
	var nuevoEvento evento.Evento
	if ID64 < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&nuevoEvento); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	services.EventoService.UpdateEvento(nuevoEvento, ID64)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&nuevoEvento); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func (c *eventoController) GetEventos(w http.ResponseWriter, r *http.Request) {
	eventossRequeridos, _ := services.EventoService.GetEventos()
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(eventossRequeridos); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func (c *eventoController) DeleteEvento(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	ID64, _ := strconv.ParseInt(id, 10, 64)
	services.EventoService.DeleteEvento(ID64)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode("Eliminado"); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
