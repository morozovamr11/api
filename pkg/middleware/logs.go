package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler { //передается что-то имеющее метод ServeHTTP(ResponseWriter, *Request) например router := http.NewServeMux()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK, //default значение
		}
		next.ServeHTTP(wrapper, r)                                               //происходит вызов router. без middleware он бы вызвался автоматически. с middleware автоматически вызывается то что в handler те Handler: middleware.Logging(router),                                            //вместо w http.ResponseWriter передали модифицированнный wrapper чтобы сохранить код
		log.Println(wrapper.StatusCode, r.Method, r.URL.Path, time.Since(start)) //!используется пакет log в log.Println
	})
}
