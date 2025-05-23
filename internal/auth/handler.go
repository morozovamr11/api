package auth

import (
	"api/configs"

	"api/pkg/reqst"
	"api/pkg/resp"
	"fmt"
	"net/http"
)

//fmt.Println(resp.Newjson())

type AuthHandlerDeps struct { //сюда изначально передается конфиг
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config //структура у которой есть методы логин и регистрация и в нее можно записать конфиг из AuthHandlerDeps
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	} //создание структуры AuthHandler чтобы использовать потом ее методы
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		body, err := reqst.HandleBody[LoginRequest](&w, req)
		if err != nil {
			return
		}
		fmt.Println(body)
		data := LoginResponse{
			Token: "123",
		}
		resp.NewJson(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := reqst.HandleBody[RegisterRequest](&w, req)
		if err != nil {
			return
		}

		// data := LoginResponse{
		// 	Token: "123",
		// }
		// resp.NewJson(w, data, 200)
		handler.AuthService.Register(body.Email, body.Password, body.Name) //в link все сделано в handler, тут вынесли в service
	}
}
