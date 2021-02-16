package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MarcosPiesskeDev/go-store-back/pkg/controller"
	"github.com/MarcosPiesskeDev/go-store-back/pkg/database"
	"github.com/MarcosPiesskeDev/go-store-back/pkg/di"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file-> ", err)
	}

	port := ":" + os.Getenv("API_PORT")
	if port == "" {
		port = ":8080"
	}

	c := &di.Container{}
	storeController := controller.NewStoreController(c.GetStoreModel())
	clientController := controller.NewClientController(c.GetClientModel())
	productController := controller.NewProductController(c.GetProductModel())

	fmt.Printf("\nServer is running on port %s ...", port)
	http.HandleFunc("/store/", storeController.InitStoreMethods)
	http.HandleFunc("/client/", clientController.InitClientMethods)
	http.HandleFunc("/product/", productController.InitProductMethods)
	log.Fatal(http.ListenAndServe(port, nil))

	db, _ := database.GetConn()
	defer db.Close()
}
