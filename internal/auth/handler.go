package auth

import (
	"api/configs"
	"api/pkg/resp"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

//fmt.Println(resp.Newjson())

type AuthHandlerDeps struct { //сюда изначально передается конфиг
	*configs.Config
}

type AuthHandler struct {
	*configs.Config //структура у которой есть методы логин и регистрация и в нее можно записать конфиг из AuthHandlerDeps
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	} //создание структуры AuthHandler чтобы использовать потом ее методы
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//read the body
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			resp.NewJson(w, err.Error(), 402)
		}
		//validation
		validate := validator.New()
		err = validate.Struct(payload)
		if err != nil {
			resp.NewJson(w, err.Error(), 402)
			return
		}
		fmt.Println(payload)
		data := LoginResponse{
			Token: "123",
		}
		resp.NewJson(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
	}
}
