package go_scratch_prod

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// Task that needs to be accomplished
type Task struct {
	gorm.Model  `json:"-"`
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	CreatedAt   time.Time `json:"created_at"`
}

func (t Task) String() string {
	s := "{"
	s += "ID: " + t.ID + ", "
	s += "Title: " + t.Title + ", "
	s += "Description: " + t.Description + ", "
	s += "IsDone: " + strconv.FormatBool(t.IsDone) + ", "
	s += "CreatedAt: " + t.CreatedAt.String()
	s += "}"

	return s
}

// TaskRepo allows persistence of tasks
type TaskRepo interface {
	Create(ctx context.Context, task Task) error
	Get(ctx context.Context, id string) (*Task, error)
	List(ctx context.Context) ([]Task, error)
	Update(ctx context.Context, task Task) error
}

// TaskService exposes features of the application
type TaskService interface {
	Create(ctx context.Context, t Task) (*Task, error)
	Update(ctx context.Context, task Task) error
	MarkDone(ctx context.Context, id string) error
	MarkPending(ctx context.Context, id string) error
	List(ctx context.Context) ([]Task, error)
	ListDone(ctx context.Context) ([]Task, error)
	ListPending(ctx context.Context) ([]Task, error)
}

// TaskHttpService exposes TaskService over an http interface
type TaskHttpService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	MarkDone(w http.ResponseWriter, r *http.Request)
	MarkPending(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	ListDone(w http.ResponseWriter, r *http.Request)
	ListPending(w http.ResponseWriter, r *http.Request)
}
