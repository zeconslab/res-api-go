package routers

import (
	"encoding/json"
	"net/http"

	"github.com/zeconslab/res-api-go/db"
	"github.com/zeconslab/res-api-go/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}
