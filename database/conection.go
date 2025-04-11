package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database

type MongoConection struct {
	User     string
	Password string
	DbName   string
	Host     string
}

func NewMongoConection(user string, password string, dbName string, host string) *MongoConection {

	return &MongoConection{
		User:     user,
		Password: password,
		DbName:   dbName,
		Host:     host,
	}

}
func (c *MongoConection) Connect() (*mongo.Database, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env: ", err)
	}

	uri := "mongodb://" + c.User + ":" + c.Password + "@" + c.Host + ":27017"

	// Criar contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conectar ao MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// Testar a conexão
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão com MongoDB estabelecida com sucesso!")

	database := client.Database(c.DbName)
	DB = database

	return DB, nil

}
