package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// Task структура
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		fmt.Println("DB open error:", err)
		return
	}
	defer db.Close()

	// Создаём таблицу если не существует
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		completed BOOLEAN DEFAULT FALSE
	)`)
	if err != nil {
		fmt.Println("Table create error:", err)
		return
	}

	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskHandler) // Для PUT/DELETE с ID

	fmt.Println("Server starting on :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}

// Обработчик для GET/POST /tasks
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTasks(w, r)
	case http.MethodPost:
		createTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Обработчик для PUT/DELETE /tasks/{id}
func taskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/tasks/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		updateTask(w, r, id)
	case http.MethodDelete:
		deleteTask(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, completed FROM tasks")
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	taskList := []Task{}
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed); err != nil {
			http.Error(w, "Scan error", http.StatusInternalServerError)
			return
		}
		taskList = append(taskList, t)
	}

	json.NewEncoder(w).Encode(taskList)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Валидация
	if task.Title == "" {
		http.Error(w, "Invalid data: title required", http.StatusBadRequest)
		return
	}

	res, err := db.Exec("INSERT INTO tasks (title, completed) VALUES (?, ?)", task.Title, task.Completed)
	if err != nil {
		http.Error(w, "DB insert error", http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	task.ID = int(id)

	json.NewEncoder(w).Encode(task)
}

func updateTask(w http.ResponseWriter, r *http.Request, id int) {
	var updated Task
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Валидация
	if updated.Title == "" {
		http.Error(w, "Invalid data: title required", http.StatusBadRequest)
		return
	}

	res, err := db.Exec("UPDATE tasks SET title = ?, completed = ? WHERE id = ?", updated.Title, updated.Completed, id)
	if err != nil {
		http.Error(w, "DB update error", http.StatusInternalServerError)
		return
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	updated.ID = id
	json.NewEncoder(w).Encode(updated)
}

func deleteTask(w http.ResponseWriter, r *http.Request, id int) {
	var deleted Task
	err := db.QueryRow("SELECT id, title, completed FROM tasks WHERE id = ?", id).Scan(&deleted.ID, &deleted.Title, &deleted.Completed)
	if err == sql.ErrNoRows {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "DB query error", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		http.Error(w, "DB delete error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(deleted)
}
