package routes

import (
	"net/http"

	"github.com/dlufy/peacekeeper/controller"
)

func HandleHttpRequests() {
	http.HandleFunc("/product", controller.AddMiddleware(controller.HandleProductRequest))
	// http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	res.Write([]byte("you are welcome to the peaceKeeper Hub, we hope you enjoy this visit"))
	// })

}
