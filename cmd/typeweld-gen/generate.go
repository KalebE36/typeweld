package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func generateCmd() *cobra.Command {
	var lang, out string
	cmd := &cobra.Command{
		Use:   "generate <file>",
		Short: "Generate code from .twld files",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			test := args[0]

			fmt.Println(test)
		},
	}

	cmd.Flags().StringVarP(&lang, "lang", "l", "go", "Target lang for gen")
	cmd.Flags().StringVarP(&out, "out", "o", "./generated", "output")

	return cmd
}
