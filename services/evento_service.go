package services

import "github.com/galileoluna/serviceEventos/domain/evento"

var (
	EventoService eventoServiceInterface = &eventoService{}
)

type eventoServiceInterface interface {
	InsertEvento(evento evento.Evento) (*evento.Evento, error)
	GetEvento(eventoID int64) (*evento.Evento, error)
	UpdateEvento(eventoActualizada evento.Evento, eventoID int64) (*evento.Evento, error)
	GetEventos() ([]evento.Evento, error)
	DeleteEvento(eventoID int64) error
}

type eventoService struct {
}

func (s *eventoService) InsertEvento(nuevoEvento evento.Evento) (*evento.Evento, error) {
	if err := nuevoEvento.Insert(); err != nil {
		return nil, err
	}
	return &nuevoEvento, nil
}

func (s *eventoService) GetEvento(eventoID int64) (*evento.Evento, error) {
	var nuevoEvento evento.Evento
	if _, err := nuevoEvento.GetEvento(eventoID); err != nil {
		return nil, err
	}
	eventoRequerida, _ := nuevoEvento.GetEvento(eventoID)
	return eventoRequerida, nil

}

func (s *eventoService) UpdateEvento(eventoActualizada evento.Evento, eventoID int64) (*evento.Evento, error) {
	if err := eventoActualizada.Update(eventoID); err != nil {
		return nil, err
	}
	return &eventoActualizada, nil

}
func (s *eventoService) GetEventos() ([]evento.Evento, error) {
	var nuevoEvento evento.Evento
	eventos, _ := nuevoEvento.GetEventos()
	return eventos, nil
}
func (s *eventoService) DeleteEvento(eventoID int64) error {
	var nuevoEvento evento.Evento
	nuevoEvento.Delete(eventoID)
	return nil

}
