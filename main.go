package main

import (
	"fmt"
	"log"
	"password-manager/models"
	"password-manager/utils"
)

func main() {
	var user models.User

	user.SetUsername("gileno")

	pass, err := utils.GeneratePassword(20, true, true, true)

	if err != nil {
		log.Fatal("Erro ao gerar senha")
	}
	user.SetPassword(pass)

	fmt.Println(user.GetPassworld())

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
