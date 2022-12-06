package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := pix.NewGerencianet(credentials)

	
	Header := map[string]interface{}{
		"x-idempotency-key" : "dt9BHlyzrb5jrFNAdfEDVpHgiOmDbVqVxd",
	},

	body := map[string]interface{} {
		"pagador": map[string]interface{} {
				"idParticipante": "00000000-0000-0000-0000-000000000000",
				"cpf" : "12345678909",
		},
		"favorecido": map[string]interface{}{
				"contaBanco": map[string]interface{}{
					"codigoBanco" : "364",
					"agencia" : "0001",
					"documento" : "11122233344",
					"nome" : "Gorbadoc Oldbuck",
					"conta" : "000000",
					"tipoConta" : "CACC"
				}
		},
		"valor": "0.01",
		"infoPagador": "Teste.",
		"idProprio": "Order_00001",
	}

	res, err := gn.ofStartPixPayment(body, Header) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}