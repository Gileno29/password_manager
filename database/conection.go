package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database

func Connect() {
	var (
		user     string
		password string
		dbName   string
		host     string
	)

	uri := "mongodb://usuario:senha@localhost:27017"
	dbName = "nome_do_banco"

	// Criar contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conectar ao MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Testar a conexão
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão com MongoDB estabelecida com sucesso!")

	// Acessar um banco de dados e coleção
	database := client.Database(dbName)
	collection := database.Collection("nome_da_colecao")

	// Agora você pode usar a variável 'collection' para operações CRUD
	fmt.Printf("Coleção %s acessada\n", collection.Name())
	DB = database

}
