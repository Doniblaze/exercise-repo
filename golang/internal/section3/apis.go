package section3

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func RunRESTfulApiExercise() {
	taskStore := NewTaskStore()
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetAllTasksHandler(w, r, taskStore)
		case http.MethodPost:
			CreateTaskHandler(w, r, taskStore)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetTaskHandler(w, r, taskStore)
		case http.MethodPut:
			UpdateTaskHandler(w, r, taskStore)
		case http.MethodDelete:
			DeleteTaskHandler(w, r, taskStore)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
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
	var id int
	_, err := fmt.Sscanf(r.URL.Path, "/tasks/%d", &id)
	return id, err == nil
}

func RespondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
