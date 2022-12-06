package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianetOf(credentials)

	res, err := gn.of_list_participants() 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}