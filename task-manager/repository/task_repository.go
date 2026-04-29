package repository

import (
	"task-manager/database"
	"task-manager/models"
)

func GetAllTasks() ([]models.Task, error) {
	rows, err := database.DB.Query("SELECT id, title, description, status, priority, created_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var t models.Task
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.Priority, &t.CreatedAt)
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func GetTaskByID(id string) (models.Task, error) {
	var t models.Task

	err := database.DB.QueryRow(
		"SELECT id, title, description, status, priority, created_at FROM tasks WHERE id = ?",
		id,
	).Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.Priority, &t.CreatedAt)

	return t, err
}

func CreateTask(t models.Task) error {
	_, err := database.DB.Exec(
		"INSERT INTO tasks (title, description, status, priority) VALUES (?, ?, ?, ?)",
		t.Title, t.Description, t.Status, t.Priority,
	)
	return err
}

func UpdateTask(id string, t models.Task) error {
	_, err := database.DB.Exec(
		"UPDATE tasks SET title=?, description=?, status=?, priority=? WHERE id=?",
		t.Title, t.Description, t.Status, t.Priority, id,
	)
	return err
}

func DeleteTask(id string) error {
	_, err := database.DB.Exec("DELETE FROM tasks WHERE id=?", id)
	return err
}
