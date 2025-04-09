package commands

import (
	"fmt"
	"log"
	"password-manager/utils"

	"github.com/spf13/cobra"
)

var comands []*cobra.Command

func LoadComands() []*cobra.Command {

	return comands
}

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

		},
	}

	cmd.Flags().StringVarP(&l, "passworld length", "l", "", "Tamanho do password")
	cmd.Flags().StringVarP(&u, "use uper letters", "u", "", "Uper case ")
	cmd.Flags().StringVarP(&s, "use especial caracters", "s", "", "Caracteres especiais")
	cmd.Flags().StringVarP(&d, "use digits", "s", "", "Digitos na senha")
	comands = append(comands, cmd)
}
