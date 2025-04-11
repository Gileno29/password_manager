package commands

import (
	"fmt"
	"log"
	"password-manager/models"
	"password-manager/utils"

	"github.com/spf13/cobra"
)

var comands []*cobra.Command

func LoadComands() []*cobra.Command {

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

			user_ := models.NewUser(pass, user, "", &models.Conta{Tipo: "sem titpo", Descricao: "sem descricao"})

		},
	}

	/*v, _ := cmd.Flags().GetInt("l")
	l = v*/

	cmd.Flags().IntVarP(&l, "passworld length", "l", 8, "Tamanho do password")
	cmd.Flags().BoolVarP(&u, "use uper letters", "u", true, "Uper case ")
	cmd.Flags().BoolVarP(&s, "use especial caracters", "s", true, "Caracteres especiais")
	cmd.Flags().BoolVarP(&d, "use digits", "s", true, "Digitos na senha")
	comands = append(comands, cmd)
}

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

	cmd.Flags().IntVarP(&l, "passworld length", "l", 8, "Tamanho do password")
	cmd.Flags().BoolVarP(&u, "use uper letters", "u", true, "Uper case ")
	cmd.Flags().BoolVarP(&s, "use especial caracters", "s", true, "Caracteres especiais")
	cmd.Flags().BoolVarP(&d, "use digits", "s", true, "Digitos na senha")
	cmd.Flags().StringVarP(&user, "create a username", "user", "", "usuario")
	comands = append(comands, cmd)

}
