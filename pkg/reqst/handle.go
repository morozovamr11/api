package reqst

import (
	"api/pkg/resp"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, req *http.Request) (*T, error) {
	body, err := Decode[T](req.Body)
	if err != nil {
		resp.NewJson(*w, err.Error(), 402)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		resp.NewJson(*w, err.Error(), 402)
		return nil, err
	}
	return &body, nil
}
