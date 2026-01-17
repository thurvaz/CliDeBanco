package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thurvaz/CliDeBanco/config"
	"github.com/thurvaz/CliDeBanco/internal"
)

func main() {

	err := config.Init()
	if err != nil {
		panic(err)
	}

	var rootCmd = &cobra.Command{
		Use:  "meuapp",
		Args: cobra.MaximumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Banco Do Thurvaz")
		},
	}
	rootCmd.AddCommand(internal.CreateAccount)
	rootCmd.Execute()
}
