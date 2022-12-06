package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials	
	gn := pix.NewGerencianet(credentials)

	const txid = "adssshdsjdsjeyccdyddsasdstxid2901"


	res, err := gn.ListDueChargesId(txid) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}