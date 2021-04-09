package controllers

import "net/http"

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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("insert"))
}
func (c *eventoController) GetEvento(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get"))
}
func (c *eventoController) UpdateEvento(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("upd"))
}
func (c *eventoController) GetEventos(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("gets"))
}
func (c *eventoController) DeleteEvento(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("dels"))
}
