package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeconslab/res-api-go/db"
	"github.com/zeconslab/res-api-go/models"
)

/*
Funcion para obtener usuarios de la base de datos
*/
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}

/*
Funcion para obtener un usuario de la base de datos
*/
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	//Referencia al objeto (Tabla DB)
	var user models.User

	//Almacenar los datos que se reciben atraves del request
	params := mux.Vars(r)

	//Buscar en la base de datos acorde a los datos que se reciben
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found!"))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

/*
Funcion para crear un usuario
*/
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	//Decodificamos lo que nos envia la peticion del body
	json.NewDecoder(r.Body).Decode(&user)
	//Creamos el usuario
	createUser := db.DB.Create(&user)
	err := createUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}

/*
Funcion para eliminar un suario
*/
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found!"))
		return
	}

	db.DB.Unscoped().Delete(&user)
	w.Write([]byte("User deleted!"))
}

/*
Funcion para actualizar un usuario
*/
func PutUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found!"))
		return
	}

	//Actualizar el usuario
	db.DB.Model(&user).Updates(models.User{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email})

	json.NewEncoder(w).Encode(&user)

}
