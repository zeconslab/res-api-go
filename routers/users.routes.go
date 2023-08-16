package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeconslab/res-api-go/db"
	"github.com/zeconslab/res-api-go/models"
	"github.com/zeconslab/res-api-go/validations"
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

	//Validar si el usuario existe
	if validations.ValidarUsuario(params["id"]) == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User not found!"))
		return
	}
	//Buscar en la base de datos acorde a los datos que se reciben
	db.DB.First(&user, params["id"])

	//Cargar las tareas del usuario
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)
}

/*
Funcion para crear un usuario
*/
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	//Decodificamos lo que nos envia la peticion del body
	json.NewDecoder(r.Body).Decode(&user)
	emailBody := user.Email

	//Validar si existe el email
	if validations.ValidarEmail(emailBody) == true {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email already exists!"))
		return
	}
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
	//Validar si el usuario existe
	if validations.ValidarUsuario(params["id"]) == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User not found!"))
		return
	}
	db.DB.Unscoped().Delete(&user)
	w.Write([]byte("User deleted!"))
}

/*
Metodo para ctualizar usuario
*/
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	//Validar si el usuario existe
	if validations.ValidarUsuario(params["id"]) == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User not found!"))
		return
	}
	db.DB.Save(&user)
	json.NewEncoder(w).Encode(&user)

}
