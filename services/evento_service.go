package services

import "github.com/galileoluna/serviceEventos/domain/evento"

var (
	EventoService eventoServiceInterface = &eventoService{}
)

type eventoServiceInterface interface {
	InsertEvento(evento evento.Evento) (*evento.Evento, error)
	GetEvento(eventoID int64) (*evento.Evento, error)
	UpdateEvento(eventoActualizada evento.Evento, eventoID int64) (*evento.Evento, error)
	GetEventos(id_comercio int64) ([]evento.Evento, error)
	DeleteEvento(eventoID int64) error
}

type eventoService struct {
}

func (s *eventoService) InsertEvento(evento evento.Evento) (*evento.Evento, error) {

}

func (s *eventoService) GetEvento(eventoID int64) (*evento.Evento, error) {

}
func (s *eventoService) UpdateEvento(eventoActualizada evento.Evento, eventoID int64) (*evento.Evento, error) {

}
func (s *eventoService) GetEventos(id_comercio int64) ([]evento.Evento, error) {

}
func (s *eventoService) DeleteEvento(eventoID int64) error {

}
