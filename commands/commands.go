package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"password-manager/database"
	"password-manager/models"
	"password-manager/repository"
	"password-manager/utils"

	"github.com/spf13/cobra"
)

var comands []*cobra.Command

func LoadComands() []*cobra.Command {
	GeneratePassword()

	return comands
}

var (
	l       int
	u, s, d bool
	user    string
)

func GeneratePassword() {

	var cmd = &cobra.Command{
		Use:   "generate",
		Short: "gera um password aletorio para um usuario",
		Run: func(cmd *cobra.Command, args []string) {
			// validations

			pass, err := utils.GeneratePassword(l, u, s, d)

			if err != nil {
				log.Fatal("Erro ao gerar senha")
			}

			fmt.Println(pass)

			user_ := models.NewUser(user, pass, "", models.NewConta("sem titpo", "sem descricao"))

			db, _ := database.NewMongoConection(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST")).Connect()

			repo := repository.NewUserRepository(db)

			repo.Create(*user_)
			defer db.Client().Disconnect(context.Background())

		},
	}

	/*v, _ := cmd.Flags().GetInt("l")
	l = v*/

	cmd.Flags().IntVarP(&l, "passworld length", "l", 8, "Tamanho do password")
	cmd.Flags().BoolVarP(&u, "use uper letters", "u", true, "Uper case ")
	cmd.Flags().BoolVarP(&s, "use especial caracters", "s", true, "Caracteres especiais")
	cmd.Flags().BoolVarP(&d, "use digits", "d", true, "Digitos na senha")
	cmd.Flags().StringVarP(&user, "user", "o", "", "Digitos na senha")
	comands = append(comands, cmd)
}

/*
func CreateUser() {

	var cmd = &cobra.Command{
		Use:   "generate-user",
		Short: "gera um password aletorio para um usuario",
		Run: func(cmd *cobra.Command, args []string) {
			// validations

			pass, err := utils.GeneratePassword(l, u, s, d)

			if err != nil {
				log.Fatal("Erro ao gerar senha")
			}

			fmt.Println(pass)

		},
	}

	/*v, _ := cmd.Flags().GetInt("l")
	l = v*/

/*	cmd.Flags().IntVarP(&l, "passworld length", "l", 8, "Tamanho do password")
	cmd.Flags().BoolVarP(&u, "use uper letters", "u", true, "Uper case ")
	cmd.Flags().BoolVarP(&s, "use especial caracters", "s", true, "Caracteres especiais")
	cmd.Flags().BoolVarP(&d, "use digits", "s", true, "Digitos na senha")
	cmd.Flags().StringVarP(&user, "create a username", "user", "", "usuario")
	comands = append(comands, cmd)

}*/
