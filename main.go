package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dlufy/peacekeeper/admin"
	"github.com/dlufy/peacekeeper/database"
	"github.com/dlufy/peacekeeper/routes"
)

func main() {
	fmt.Println("it's peace Keeper's Hub / Heart")
	admin.Initialize("./admin/peacekeeper.ini")
	//import route handler
	routes.HandleHttpRequests()

	err := database.Connect()
	if err != nil {
		log.Fatal("error initliazing the DB", err)
	}
	mainConfig := admin.GetConfig()
	fmt.Println("server is listening at port ", mainConfig.Server.Port)
	if err := http.ListenAndServe(":"+mainConfig.Server.Port, nil); err != nil {
		fmt.Println("not able to start the server", err)
		return
	}

}
