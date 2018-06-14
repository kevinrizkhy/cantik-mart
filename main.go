package main

import (
	"fmt"
	"github.com/gorilla/mux"
	config "github.com/wellcode/pardev/proposal/config"
	router "github.com/wellcode/pardev/proposal/controller/router"
	"net/http"
	"os"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		addr = config.Base_Port
	}
	r := mux.NewRouter()
	r.HandleFunc("/", router.Sign)
	r.HandleFunc("/home", router.Home)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.Handle("/assets/", r)
	http.ListenAndServe(addr, r)
}
