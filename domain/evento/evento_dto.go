package evento

type Evento struct {
	ID_Evento          int64
	Nombre             string
	Descripcion        string
	HoraDeInicio       string
	HoraDeFinalizacion string
}

func NewEvento(id int64, nombre string, descripcion string, horaInicio string, horaFinal string) Evento {
	return Evento{
		ID_Evento:          id,
		Nombre:             nombre,
		Descripcion:        descripcion,
		HoraDeInicio:       horaInicio,
		HoraDeFinalizacion: horaFinal,
	}
}
