package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"password-manager/database"
)

func main() {

	mongo := database.NewMongoConection(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"))

	conection, err := mongo.Connect()

	if err != nil {
		log.Fatal("Erro ao se conectar no banco de dados ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer conection.Client().Disconnect(ctx)
	err = conection.CreateCollection(ctx, "teste")

	if err != nil {
		log.Fatal("Erro ao criar a colection", err)
	}

	collection := conection.Collection("teste")
	fmt.Println("Colecao", collection.Database().Name())

	/*privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}*/

	/*publicKey := &privateKey.PublicKey

	encrypted, err := utils.EncryptPasswordWithKey(user.GetPassworld(), publicKey)
	if err != nil {
		panic(err)
	}

	fmt.Println(encrypted)

	// Para verificar (em outro momento):
	// 1. Primeiro descriptografar com a chave privada
	bcryptHash, err := utils.DecryptPasswordWithKey(encrypted, privateKey)
	if err != nil {
		panic(err)
	}

	fmt.Println(bcryptHash)

	// 2. Depois verificar com bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(bcryptHash), []byte(user.GetPassworld()))
	if err != nil {
		println("Senha inválida!")
	} else {
		println("Senha válida!")
	}*/

}
