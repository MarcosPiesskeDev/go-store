package routes

import (
	"go-store-back/src/controller"
	"log"
	"net/http"
)

func ConfigRoutes() {
	http.HandleFunc("/store/", controller.InitStoreMethods)
	http.HandleFunc("/client/", controller.InitClientMethods)
	http.HandleFunc("/product/", controller.InitProductMethods)
	log.Fatal(http.ListenAndServe(":9000", nil))

}
