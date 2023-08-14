package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeconslab/res-api-go/db"
	"github.com/zeconslab/res-api-go/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	if len(tasks) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Tasks not found!"))
		return
	}

	json.NewEncoder(w).Encode(&tasks)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	db.DB.Create(&task)
	json.NewEncoder(w).Encode(&task)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	db.DB.First(&task, r.URL.Query().Get("id"))
	json.NewEncoder(w).Encode(&task)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	db.DB.First(&task, r.URL.Query().Get("id"))
	json.NewDecoder(r.Body).Decode(&task)
	db.DB.Save(&task)
	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found!"))
		return
	}

	db.DB.Unscoped().Delete(&task)
	json.NewEncoder(w).Encode("Task deleted!")
}
