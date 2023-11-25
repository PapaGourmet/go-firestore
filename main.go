package main

import (
	"fmt"

	irepositorys "github.com/PapaGourmet/estudosGO/irepositorys"
	services "github.com/PapaGourmet/estudosGO/services"
)



func main() {
	_subscriber, error := irepositorys.GetAllSubscribers(services.FisrestoreService{})

	if error != nil {
		fmt.Println(error.Error())
	}

	fmt.Println(_subscriber)
}
