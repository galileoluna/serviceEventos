package main

import (
	"fmt"

	"github.com/galileoluna/serviceEventos/datasources"
	"github.com/galileoluna/serviceEventos/domain/evento"
	"github.com/galileoluna/serviceEventos/services"
)

func Main() {
	datasources.Init()
	conciertoMona := evento.NewEvento(1, "Cold play", "Unico en el estadio la plata", "06-09-1997", "06-09-1998")

	conciertoMonaAct := evento.NewEvento(1, "Concierto Mega Death", "Unico en el estadio Kempes", "06-09-1997", "06-09-1998")
	//conciertoMona.Insert()

	services.EventoService.InsertEvento(conciertoMona)
	services.EventoService.UpdateEvento(conciertoMonaAct, 1)

	fmt.Println(services.EventoService.GetEvento(3))
	fmt.Println(services.EventoService.GetEventos())

}
