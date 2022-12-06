package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianetOf(credentials)

	body := map[string]interface{} {
		"redirectURL" : "https://your-domain.com.br/redirect/",
		"webhookURL" : "https://your-domain.com.br/webhook/",
	}

	res, err := gn.ofConfigUpdate(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}