package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/marcos-dev88/go-store-back/pkg/controller"
	"github.com/marcos-dev88/go-store-back/pkg/database"
	"github.com/marcos-dev88/go-store-back/pkg/di"
)

func init() {
	if err := defineEnvs(".env"); err != nil {
		log.Printf("Error to load .env -> %v", err)
	}
}

func main() {

	port := ":" + os.Getenv("API_PORT")
	if port == "" {
		port = ":8080"
	}

	c := &di.Container{}
	storeController := controller.NewStoreController(c.GetStoreModel())
	clientController := controller.NewClientController(c.GetClientModel())
	productController := controller.NewProductController(c.GetProductModel())

	db, _ := database.GetConn()
	defer db.Close()

	fmt.Printf("\nServer is running on port %s ...", port)
	http.HandleFunc("/store/", storeController.InitStoreMethods)
	http.HandleFunc("/client/", clientController.InitClientMethods)
	http.HandleFunc("/product/", productController.InitProductMethods)
	log.Fatal(http.ListenAndServe(port, nil))
}

func defineEnvs(fileName string) error {

	envs := make(map[string]string)

	file, err := os.Open(fileName)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("We got an error -> %v", err)
		}
	}(file)

	if err != nil {
		return err
	}

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		envSplit := strings.SplitN(sc.Text(), "=", 2)
		if len(envSplit) > 1 {
			envs[envSplit[0]] = envSplit[1]
		}
	}

	for key, value := range envs {
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	return nil
}