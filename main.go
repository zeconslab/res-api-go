package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeconslab/res-api-go/db"
	"github.com/zeconslab/res-api-go/models"
	"github.com/zeconslab/res-api-go/routers"
)

func main() {

	//Ejecutar conexion con la base de datos
	db.DBconection()

	//Migracion de la base de datos
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	//Creacion de router
	route := mux.NewRouter()

	//Listado de routers(Paginas)
	route.HandleFunc("/", routers.HomeHandler)
	route.HandleFunc("/users", routers.GetUsersHandler).Methods("GET")
	route.HandleFunc("/user/{id}", routers.GetUserHandler).Methods("GET")
	route.HandleFunc("/users", routers.PostUserHandler).Methods("POST")
	route.HandleFunc("/user/{id}", routers.DeleteUserHandler).Methods("DELETE")
	route.HandleFunc("/user/{id}", routers.PutUserHandler).Methods("PUT")

	//Inicializar servidor de escucha
	port := "4000"
	hostname := "localhost"
	log.Println(port)
	log.Print("Server listen port ", port, ". Go to http://", hostname, ":", port, "/")
	http.ListenAndServe(":"+port, route)

}
