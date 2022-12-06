package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := pix.NewGerencianet(credentials)

	const id = "1e33f858-f2c1-49e0-a161-7f8157da27e1"

	res, err := gn.DetailReport(id) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}