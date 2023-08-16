package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/zeconslab/res-api-go/db"
	"github.com/zeconslab/res-api-go/routers"
)

func main() {
	//Cargar variavles .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Ejecutar conexion con la base de datos
	db.DBconection(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

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

	//paginas de tareas
	route.HandleFunc("/tasks", routers.GetTasksHandler).Methods("GET")
	route.HandleFunc("/task/{id}", routers.GetTaskHandler).Methods("GET")
	route.HandleFunc("/task", routers.PostTaskHandler).Methods("POST")
	route.HandleFunc("/task/{id}", routers.DeleteTaskHandler).Methods("DELETE")
	route.HandleFunc("/task/{id}", routers.UpdateDTaskHandler).Methods("PUT")

	//Inicializar servidor de escucha
	port := "3000"
	hostname := "localhost"
	log.Print("Server listen port ", port, ". Go to http://", hostname, ":", port, "/")
	http.ListenAndServe(":"+port, route)

}
