package evento

import (
	"fmt"

	"github.com/galileoluna/serviceEventos/datasources"
)

func Test() {
	datasources.Init()
	conciertoMona := NewEvento(1, "Concierto Mona Jimenez", "Unico en el estadio Kempes", "06-09-1997", "06-09-1998")

	conciertoMona.Insert()

	fmt.Println(conciertoMona.GetEvento(1))
	fmt.Println(conciertoMona.GetEventos())

}
