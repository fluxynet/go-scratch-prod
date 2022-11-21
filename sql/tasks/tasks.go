package Tasks

import (
	"context"
	"database/sql"

	gsp "github.com/fluxynet/go-scratch-prod"
)

var _ gsp.TaskRepo = Repo{}

func New(db *sql.DB) Repo {
	return Repo{db: db}
}

type Repo struct {
	db *sql.DB
}

func (r Repo) Create(ctx context.Context, task gsp.Task) error {
	query := `INSERT INTO tasks (id, title, description, is_done, created_at) VALUES (?,?,?,?,?)`

	_, err := r.db.ExecContext(
		ctx,
		query,
		task.ID,
		task.Title,
		task.Description,
		task.IsDone,
		task.CreatedAt,
	)

	return err
}

func (r Repo) Get(ctx context.Context, id string) (*gsp.Task, error) {
	query := `SELECT id, title, description, is_done, created_at FROM tasks WHERE id = ?`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, gsp.ErrEntityNotFound
	}

	var task gsp.Task
	err = rows.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.IsDone,
		&task.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r Repo) List(ctx context.Context) ([]gsp.Task, error) {
	query := `SELECT id, title, description, is_done, created_at FROM tasks`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var tasks []gsp.Task

	for rows.Next() {
		var task gsp.Task
		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.IsDone,
			&task.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r Repo) Update(ctx context.Context, task gsp.Task) error {
	query := `UPDATE tasks SET title=?, description=?, is_done=?, created_at=? WHERE id = ?`

	_, err := r.db.ExecContext(
		ctx,
		query,
		task.Title,
		task.Description,
		task.IsDone,
		task.CreatedAt,
		task.ID,
	)

	return err
}
