package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MarcosPiesskeDev/go-store-back/pkg/config"
	"github.com/MarcosPiesskeDev/go-store-back/pkg/controller"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file-> ", err)
	}

	port := os.Getenv("API_PORT")
	if port == "" {
		port = ":8080"
	}

	fmt.Printf("\nServer is running on port %s ...", port)
	http.HandleFunc("/store/", controller.InitStoreMethods)
	http.HandleFunc("/client/", controller.InitClientMethods)
	http.HandleFunc("/product/", controller.InitProductMethods)
	log.Fatal(http.ListenAndServe(port, nil))

	db, _ := config.GetDb()
	defer db.Close()
}
