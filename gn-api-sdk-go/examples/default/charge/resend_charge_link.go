package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	res, err := gn.ResendChargeLink(1) //no lugar do 1 informe o id da cobrança desejada

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}