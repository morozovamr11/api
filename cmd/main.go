package main

import (
	//"api/configs"
	"api/configs"
	"api/internal/auth"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig() //загрузили конфиг
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	//http.ListenAndServe(":8081", nil) //используется дефолтный server max
}
