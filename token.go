package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

func getToken(ask bool) string {
	token := ""
	if !ask {
		token = os.Getenv("ghdefaultbranchToken")
	}
	if token == "" {
		fmt.Print("Enter Github Access Token (must have repo admin permissions): ")
		bytePassword, err := term.ReadPassword(0)
		fmt.Println()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		token = string(bytePassword)
	}
	return token
}
