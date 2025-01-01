package main

import (
	"fmt"

	httpcustomer "github.com/renatospaka/poc-http/customer/http"
	httpproduct "github.com/renatospaka/poc-http/product/http"
	"github.com/renatospaka/poc-http/server"
)

func main() {
	fmt.Println("abrindo o serviço")

	// Cria o servidor HTTP e obtém o canal de prontidão
	mux, serverReady := server.NewServer(":8080")

	// Aguarda até que o servidor esteja pronto
	<-serverReady

	// Registra as rotas para customers e products
	httpcustomer.RegisterRoutes(mux)
	httpproduct.RegisterRoutes(mux)

	fmt.Println("Rotas registradas e servidor está ouvindo em", mux.Addr())

	// Mantém o servidor rodando
	select {}
}
