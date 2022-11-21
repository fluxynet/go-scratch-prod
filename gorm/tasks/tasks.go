package tasks

import (
	"context"
	"log"

	"gorm.io/gorm"

	gsp "github.com/fluxynet/go-scratch-prod"
)

var _ gsp.TaskRepo = Repo{}

func New(db *gorm.DB) Repo {
	return Repo{db: db}
}

type Repo struct {
	db *gorm.DB
}

func (r Repo) Create(ctx context.Context, task gsp.Task) error {
	err := r.db.Create(&task).Error

	log.Println("Created task (err=", err, ") ", task)

	return err
}

func (r Repo) Get(ctx context.Context, id string) (*gsp.Task, error) {
	var task gsp.Task

	err := r.db.First(&task, "id = ?", id).Error
	log.Println("Got task", task)

	return &task, err
}

func (r Repo) List(ctx context.Context) ([]gsp.Task, error) {
	var tasks []gsp.Task

	err := r.db.Find(&tasks).Error
	log.Println("Got tasks", tasks)

	return tasks, err
}

func (r Repo) Update(ctx context.Context, task gsp.Task) error {
	err := r.db.Save(&task).Error

	log.Println("Updated task (err=", err, ") ", task)

	return err
}
