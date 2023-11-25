package connection

import (
	"context"
	"log"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Conn() (*firestore.Client, context.Context, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("key.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, error := app.Firestore(ctx)
	if error != nil {
		log.Fatalln(err)
	} 
	
	return client, ctx, error	
}