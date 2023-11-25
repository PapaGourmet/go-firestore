package services

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	connection "github.com/PapaGourmet/estudosGO/firestore"
	irepositorys "github.com/PapaGourmet/estudosGO/irepositorys"
	"google.golang.org/api/iterator"
)

type PersonError struct {
	Message string
}

func (e PersonError) Error() string {
	return fmt.Sprintf("Ocorreu um erro: %s", e.Message)
}

type FisrestoreService struct{}

func multipleDocs(ctx context.Context, client *firestore.Client) (irepositorys.Subscriber, error) {

	init := irepositorys.Subscriber{
		Last_update: 0,
		Name:        "",
		Phone:       "",
		PortNumber:  "",
		Session_id:  "",
		Status:      false,
		Cep: irepositorys.PostalCode{
			District:   "",
			PostalCode: "",
			City:       "",
			Street:     "",
			State:      ""}}

	iter := client.Collection("subscribers").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			println(err.Error())
			return init, err
		}

		cepMap, found := doc.Data()["cep"].(map[string]interface{})
		if found {

			City := cepMap["city"].(string)
			District := cepMap["district"].(string)
			Postalcode := cepMap["postalcode"].(string)
			State := cepMap["state"].(string)
			Street := cepMap["street"].(string)

			_cep := irepositorys.PostalCode{
				City:       City,
				District:   District,
				PostalCode: Postalcode,
				State:      State,
				Street:     Street,
			}


			Last_update := doc.Data()["last_update"].(int64)
			Name := doc.Data()["name"].(string)
			Phone := doc.Data()["phone"].(string)
			PortNumber := doc.Data()["portnumber"].(string)
			Session_id := doc.Data()["session_id"].(string)
			Status := doc.Data()["status"].(bool)

			_subscriber := irepositorys.Subscriber{
				Last_update: Last_update,
				Name:        Name,
				Phone:       Phone,
				PortNumber:  PortNumber,
				Session_id:  Session_id,
				Status:      Status,
				Cep:         _cep,
			}

			return _subscriber, nil

		} else {
			return init, PersonError{Message: "Erro de chave e valor"}
		}
	}

	return init, nil
}

func (f FisrestoreService) GetAll() (irepositorys.Subscriber, error) {
	client, ctx, error := connection.Conn()

	if error != nil {
		println(error.Error())
	}

	_subscriber, error := multipleDocs(ctx, client)

	defer client.Close()

	return _subscriber, error

}
