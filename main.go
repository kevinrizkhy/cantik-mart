package main

import (
	"fmt"
	config "github.com/pardev/cantik-mart/config"
	router "github.com/pardev/cantik-mart/controller/router"
	database "github.com/pardev/cantik-mart/model/db"
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
	fmt.Println("Listening on port -> " + addr)
	fmt.Println("Connecting to Database...")
	database.Connect()
	fmt.Println("Database Connected.")
	fmt.Println("Listening....")
	router.Routes(addr)

}
