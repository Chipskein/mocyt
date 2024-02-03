package youtube

import (
	"log"
	"net/http"
	"time"
)

var server *http.Server

func handlerfunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		code := r.URL.Query().Get("code")
		token, err := ExchangeCode(code)
		if err != nil {
			log.Println(err)
		}
		saveToken("token.json", token)
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
func InitServer() {
	http.HandleFunc("/", handlerfunc)
	server = &http.Server{
		Addr:    ":5000",
		Handler: http.DefaultServeMux,
	}
	log.Println("Iniciando Servidor na porta 5000")
	log.Fatal(server.ListenAndServe())
}

func KillServer() {
	KillChannel <- true
	server.Close()
}
