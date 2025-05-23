package main

import (
	"password-manager/commands"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "etl",
	Short: "Uma aplicação CLI para registro de passwords",
}

func main() {

	comands := commands.LoadComands()
	for i := 0; i < len(comands); i++ {
		rootCommand.AddCommand(comands[i])

	}
	rootCommand.Execute()

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
