package youtube

import (
	"log"
	"net/http"
	"time"
)

var server *http.Server
var token_path string

func handlerfunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		code := r.URL.Query().Get("code")
		token, err := ExchangeCode(code)
		if err != nil {
			log.Println(err)
		}
		saveToken(token_path, token)
		log.Println("Killing Server")
		go func() {
			time.Sleep(10 * time.Second)
			KillServer()
		}()
		w.Write([]byte("<html><h1>Success Close this page the server will be Killed</h1></html>"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func InitServer(token_json_path string) {
	token_path = token_json_path
	http.HandleFunc("/", handlerfunc)
	server = &http.Server{
		Addr:    ":5000",
		Handler: http.DefaultServeMux,
	}
	log.Println("Iniciando Servidor na porta 5000")
	server.ListenAndServe()
}

func KillServer() {
	KillChannel <- true
	server.Close()
}
