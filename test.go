package main

import (
	"fmt"

	"github.com/galileoluna/serviceEventos/datasources"
	"github.com/galileoluna/serviceEventos/domain/evento"
)

func main() {
	datasources.Init()
	conciertoMona := evento.NewEvento(1, "Concierto Mona Jimenez", "Unico en el estadio Kempes", "06-09-1997", "06-09-1998")

	conciertoMona.Insert()

	fmt.Println(conciertoMona.GetEvento(2))
	fmt.Println(conciertoMona.GetEventos())

}
