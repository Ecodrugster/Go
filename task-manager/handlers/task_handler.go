package handlers

import (
	"encoding/json"
	"net/http"
	"task-manager/models"
	"task-manager/repository"
	"task-manager/utils"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := repository.GetAllTasks()
	if err != nil {
		utils.JSON(w, 500, err.Error())
		return
	}
	utils.JSON(w, 200, tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	task, err := repository.GetTaskByID(id)
	if err != nil {
		utils.JSON(w, 404, "Task not found")
		return
	}
	utils.JSON(w, 200, task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	if task.Title == "" {
		utils.JSON(w, 400, "Title required")
		return
	}

	repository.CreateTask(task)
	utils.JSON(w, 201, "Created")
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	err := repository.UpdateTask(id, task)
	if err != nil {
		utils.JSON(w, 404, "Not found")
		return
	}

	utils.JSON(w, 200, "Updated")
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := repository.DeleteTask(id)
	if err != nil {
		utils.JSON(w, 404, "Not found")
		return
	}

	utils.JSON(w, 200, "Deleted")
}