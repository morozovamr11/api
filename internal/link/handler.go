package link

import (
	"fmt"
	"net/http"
)

type LinkHandlerDeps struct { //сюда изначально передается конфиг

}
type LinkHandler struct {
	//структура у которой есть методы логин и регистрация и в нее можно записать конфиг из AuthHandlerDeps
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{} //создание структуры LinkHandler чтобы использовать потом ее методы
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("GET /{alias}", handler.GoTo())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Create")
	}
}
func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("GoTo")
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Update")
	}
}
func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Delete")
	}
}
