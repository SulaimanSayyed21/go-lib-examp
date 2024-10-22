package main

import (
	"fmt"
	"os"

	"github.com/SulaimanSayyed21/go-lib-examp/modc"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{ // Capitalize 'Command'
		Use:   "greet",
		Short: "Greets from all modules",                              // Capitalize 'Short'
		Long:  "This command greets module A, Module B, and Module C", // Fix typo and capitalize
		Run: func(cmd *cobra.Command, args []string) { // Capitalize 'Command'
			fmt.Println(modc.SayHelloC()) // Change 'println' to 'Println'
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
