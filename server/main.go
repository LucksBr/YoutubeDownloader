package main

import (
	"fmt"
	"server/resource/router"
	"net/http"
)

func main(){
	newRouter := router.CreateRouter()
	
	newRouter.RegisterRoute(router.GET, "/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Olá! Esta é uma rota GET.")
	})


	http.ListenAndServe(":8080", newRouter.GetServerHttp())
}