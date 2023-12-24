package section3

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

func RunRESTfulApiExercise() {

	r := mux.NewRouter() // Create an instance of the router
	taskStore := NewTaskStore()

	// Register the route for GET requests to /get-all-tasks
	r.HandleFunc("/get-all-tasks", func(w http.ResponseWriter, r *http.Request) {
		GetAllTasksHandler(w, r, taskStore)
	}).Methods("GET")
	// Regiser the route for POST requests to /create-tasks
	r.HandleFunc("/create-tasks", func(w http.ResponseWriter, r *http.Request) {
		CreateTaskHandler(w, r, taskStore)
	}).Methods("POST")
	// Register the route for GET requests to /get-tasks
	r.HandleFunc("/get-tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetTaskHandler(w, r, taskStore)
	}).Methods("GET")
	// Register the route for PUT requests to /update-tasks
	r.HandleFunc("/update-tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		UpdateTaskHandler(w, r, taskStore)
	}).Methods("PUT")
	// Register the route for DELETE requests to /delete-tasks
	r.HandleFunc("/delete-tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		DeleteTaskHandler(w, r, taskStore)
	}).Methods("DELETE")

	http.Handle("/", r)
	// Start the HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// model
// Task represents a todo task.
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// TaskStore is an in-memory storage for tasks.
type TaskStore struct {
	Tasks   map[int]Task
	Counter int
	Mu      sync.Mutex
}

// service class
// NewTaskStore creates a new TaskStore.
func NewTaskStore() *TaskStore {
	return &TaskStore{
		Tasks: make(map[int]Task),
	}
}

// CreateTask creates a new task.
func (ts *TaskStore) CreateTask(title string) Task {
	ts.Mu.Lock()
	defer ts.Mu.Unlock()

	ts.Counter++
	task := Task{
		ID:    ts.Counter,
		Title: title,
	}
	ts.Tasks[task.ID] = task
	return task
}

// GetTask retrieves a task by ID.
func (ts *TaskStore) GetTask(id int) (Task, bool) {
	ts.Mu.Lock()
	defer ts.Mu.Unlock()

	task, exists := ts.Tasks[id]
	return task, exists
}

// GetAllTasks retrieves all tasks.
func (ts *TaskStore) GetAllTasks() []Task {
	ts.Mu.Lock()
	defer ts.Mu.Unlock()

	var tasks []Task
	for _, task := range ts.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// UpdateTask updates an existing task.
func (ts *TaskStore) UpdateTask(id int, title string) bool {
	ts.Mu.Lock()
	defer ts.Mu.Unlock()

	if _, exists := ts.Tasks[id]; exists {
		ts.Tasks[id] = Task{ID: id, Title: title}
		return true
	}
	return false
}

// DeleteTask deletes a task by ID.
func (ts *TaskStore) DeleteTask(id int) bool {
	ts.Mu.Lock()
	defer ts.Mu.Unlock()

	if _, exists := ts.Tasks[id]; exists {
		delete(ts.Tasks, id)
		return true
	}
	return false
}

// controllers class
func GetAllTasksHandler(w http.ResponseWriter, r *http.Request, taskStore *TaskStore) {
	tasks := taskStore.GetAllTasks()
	RespondJSON(w, tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request, taskStore *TaskStore) {
	var input struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task := taskStore.CreateTask(input.Title)
	RespondJSON(w, task)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request, taskStore *TaskStore) {

	fmt.Println("Received request path:", r.URL.RequestURI())
	id, ok := GetIDFromURL(r)
	if !ok {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, exists := taskStore.GetTask(id)
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	RespondJSON(w, task)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request, taskStore *TaskStore) {
	id, ok := GetIDFromURL(r)
	if !ok {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var input struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !taskStore.UpdateTask(id, input.Title) {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request, taskStore *TaskStore) {
	id, ok := GetIDFromURL(r)
	if !ok {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	if !taskStore.DeleteTask(id) {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetIDFromURL(r *http.Request) (int, bool) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		return 0, false
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, false
	}

	return id, true
}

func RespondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
