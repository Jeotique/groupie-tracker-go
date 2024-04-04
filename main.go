package main

import (
	"fmt"
	"log"
	"project/functions"
	"project/routes"
	"project/templates"
	"project/variables"
)

func main() {
	var err error
	variables.AccessToken, err = functions.TokenAccess()
	if err != nil {
		log.Fatalf("Failed to retrieve access token: %v", err)
	}
	data, errData := functions.GetMusics("4VL5XwfATZuAVTW471Wpro", variables.AccessToken)
	if errData != nil {
		fmt.Println(errData)
	}
	fmt.Println(data)

	templates.InitTemplates()
	routes.InitRoutes()
}
