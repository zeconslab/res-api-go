package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeconslab/res-api-go/db"
	"github.com/zeconslab/res-api-go/routers"
)

func main() {

	//Ejecutar conexion con la base de datos
	db.DBconection()

	//Creacion de router
	route := mux.NewRouter()

	/*Listado de routers(Paginas)*/
	//Pagina principal
	route.HandleFunc("/", routers.HomeHandler)
	//Pagina de usuarios
	route.HandleFunc("/users", routers.GetUsersHandler).Methods("GET")
	route.HandleFunc("/user/{id}", routers.GetUserHandler).Methods("GET")
	route.HandleFunc("/user", routers.PostUserHandler).Methods("POST")
	route.HandleFunc("/user/{id}", routers.DeleteUserHandler).Methods("DELETE")
	route.HandleFunc("/user/{id}", routers.UpdateUserHandler).Methods("PUT")

	//Pagina de tareas
	route.HandleFunc("/tasks", routers.GetTasksHandler).Methods("GET")
	route.HandleFunc("/task/{id}", routers.GetTaskHandler).Methods("GET")
	route.HandleFunc("/task", routers.PostTaskHandler).Methods("POST")
	route.HandleFunc("/task/{id}", routers.DeleteTaskHandler).Methods("DELETE")
	route.HandleFunc("/task/{id}", routers.UpdateTaskHandler).Methods("PUT")

	//Inicializar servidor de escucha
	port := "3000"
	hostname := "localhost"
	log.Print("Server listen port ", port, ". Go to http://", hostname, ":", port, "/")
	http.ListenAndServe(":"+port, route)

}
