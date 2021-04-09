package evento

import (
	"errors"

	"github.com/galileoluna/serviceEventos/datasources"
)

const (
	queryInsertEvento  = "INSERT INTO eventos(Nombre,Descripcion,HoraDeInicio,HoraDeFinalizacion) VALUES($1,$2,$3,$4)"
	queryUpdateEvento  = "UPDATE eventos SET Nombre=$1, Descripcion=$2, HoraDeInicio=$3,HoraDeFinalizacion=$4 "
	queryDeleteEvento  = "DELETE FROM eventos WHERE ID_Evento=$1; "
	queryGetEventoByID = "SELECT * FROM eventos where ID_Evento=$1;"
	queryGetEventos    = "SELECT * FROM eventos;"
)

func (nuevoEvento *Evento) Insert() error {
	_, err := datasources.Db.Exec(queryInsertEvento, nuevoEvento.Nombre, nuevoEvento.Descripcion, nuevoEvento.HoraDeInicio, nuevoEvento.HoraDeFinalizacion)
	if err != nil {
		return err
	}
	return nil
}

func (eventoActualizado *Evento) Update(id_evento int64) error {

	_, err := datasources.Db.Exec(queryUpdateEvento, id_evento, eventoActualizado.Nombre, eventoActualizado.Descripcion, eventoActualizado.HoraDeInicio, eventoActualizado.HoraDeFinalizacion)
	if err != nil {
		return err
	}
	return nil
}

func (eventoBorrado *Evento) Delete(id_evento int64) error {
	_, err := datasources.Db.Exec(queryDeleteEvento, id_evento)
	if err != nil {
		return err
	}
	return nil
}

func (eventoBuscado *Evento) GetEvento(id_evento int64) (*Evento, error) {
	var encontrada Evento
	rows, err := datasources.Db.Query(queryGetEventoByID, id_evento)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id_evento int64
		var nombre, descripcion, horarioInicio, horarioFinalizacion string

		err := rows.Scan(&id_evento, &nombre, &descripcion, &horarioInicio, &horarioFinalizacion)
		if err != nil {
			err = errors.New("Ingrese los parametros adecuados")
			return nil, err
		}
		encontrada = NewEvento(id_evento, nombre, descripcion, horarioInicio, horarioFinalizacion)
	}

	return &encontrada, nil

}

func (eventosBuscados *Evento) GetEventos() ([]Evento, error) {
	Eventos := []Evento{}

	rows, err := datasources.Db.Query(queryGetEventos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id_evento int64
		var nombre, descripcion, horarioInicio, horarioFinalizacion string
		var encontrada Evento

		err := rows.Scan(&id_evento, &nombre, &descripcion, &horarioInicio, &horarioFinalizacion)
		if err != nil {
			err = errors.New("Ingrese los parametros adecuados")
			return nil, err
		}
		encontrada = NewEvento(id_evento, nombre, descripcion, horarioInicio, horarioFinalizacion)
		Eventos = append(Eventos, encontrada)
	}
	return Eventos, err
}
