package main

import (
	"api/internal/hello"
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	hello.NewHelloHandler(router)
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
