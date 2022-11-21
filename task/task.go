package task

import (
	"context"
	"log"
	"time"

	gsp "github.com/fluxynet/go-scratch-prod"

	"github.com/google/uuid"
)

var _ gsp.TaskService = Service{}

func New(repo gsp.TaskRepo) Service {
	return Service{repo: repo}
}

type Service struct {
	repo gsp.TaskRepo
}

func (svc Service) Create(ctx context.Context, task gsp.Task) (*gsp.Task, error) {
	task.ID = uuid.NewString()
	task.CreatedAt = time.Now()

	err := svc.repo.Create(ctx, task)

	log.Println("task created", task)

	return &task, err
}

func (svc Service) Update(ctx context.Context, task gsp.Task) error {
	return svc.repo.Update(ctx, task)
}

func (svc Service) MarkDone(ctx context.Context, id string) error {
	task, err := svc.repo.Get(ctx, id)
	if err != nil {
		return err
	}

	task.IsDone = true
	return svc.repo.Update(ctx, *task)
}

func (svc Service) MarkPending(ctx context.Context, id string) error {
	task, err := svc.repo.Get(ctx, id)
	if err != nil {
		return err
	}

	task.IsDone = false
	return svc.repo.Update(ctx, *task)
}

func (svc Service) List(ctx context.Context) ([]gsp.Task, error) {
	return svc.repo.List(ctx)
}

func (svc Service) ListDone(ctx context.Context) ([]gsp.Task, error) {
	return nil, gsp.ErrNotImplemented
}

func (svc Service) ListPending(ctx context.Context) ([]gsp.Task, error) {
	return nil, gsp.ErrNotImplemented
}
