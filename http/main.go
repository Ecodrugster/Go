package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)


type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}


var students = []Student{
	{ID: 1, Name: "Ali", Age: 18},
	{ID: 2, Name: "Aruzhan", Age: 19},
	{ID: 3, Name: "Nurlan", Age: 17},
}

// 1

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Добро пожаловать в Student API")
}

// 2

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Это учебный сервер на Go")
}

// 3

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Гость"
	}

	fmt.Fprintf(w, "Привет, %s!\n", name)
}

// 4

func studentsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(students)

	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")

		var newStudent Student

		err := json.NewDecoder(r.Body).Decode(&newStudent)
		if err != nil {
			http.Error(w, "Неверный JSON", http.StatusBadRequest)
			return
		}

		if newStudent.Name == "" {
			http.Error(w, "Имя не может быть пустым", http.StatusBadRequest)
			return
		}

		if newStudent.Age <= 0 {
			http.Error(w, "Возраст должен быть больше 0", http.StatusBadRequest)
			return
		}

		students = append(students, newStudent)

		response := map[string]interface{}{
			"message": "Студент успешно добавлен",
			"student": newStudent,
		}

		json.NewEncoder(w).Encode(response)

	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}


// 5


func studentHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "id не передан", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id должен быть числом", http.StatusBadRequest)
		return
	}

	for _, s := range students {
		if s.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(s)
			return
		}
	}

	http.Error(w, "Студент не найден", http.StatusNotFound)
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/students", studentsHandler)
	http.HandleFunc("/student", studentHandler)

	fmt.Println("Server started at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска:", err)
	}
}