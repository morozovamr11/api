package main

import (
	//"api/configs"
	"api/configs"
	"api/internal/auth"
	"api/internal/link"
	"api/pkg/db"
	"api/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig() //загрузили конфиг
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//repositories
	linkRepository := link.NewLinkRepository(db)
	//handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})
	//end handlers
	//middlewarws
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router), //использование middleware, без него просто router
	}
	fmt.Println("Server is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	//http.ListenAndServe(":8081", nil) //используется дефолтный server max
}
