package main

import (
	"fmt"

	"github.com/PapaGourmet/estudosGO/internal/irepositorys"
)


func main() {
	_subscriber, error := irepositorys.GetAllSubscribers(services.FisrestoreService{})

	if error != nil {
		fmt.Println(error.Error())
	}

	fmt.Println(_subscriber)
}
