package internal

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thurvaz/CliDeBanco/config"
)

var db = config.GetSQLite()

var CreateAccount = &cobra.Command{
	Use:  "createaccount [name] [cpf] [password]",
	Args: cobra.MinimumNArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string
		var cpf string
		var password string
		if len(args) == 3 {
			name = args[0]
			cpf = args[1]
			password = args[2]
		}
		fmt.Printf("Conta criada %s %s %s", name, cpf, password)

		return nil
	},
}
