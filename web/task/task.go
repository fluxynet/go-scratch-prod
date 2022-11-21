package task

import (
	"net/http"

	gsp "github.com/fluxynet/go-scratch-prod"
	"github.com/fluxynet/go-scratch-prod/web"
)

var _ gsp.TaskHttpService = Service{}

func New(tasks gsp.TaskService) Service {
	return Service{tasks: tasks}
}

type Service struct {
	tasks gsp.TaskService
}

func (svc Service) Create(w http.ResponseWriter, r *http.Request) {
	var task gsp.Task

	if err := web.ReadJsonBodyInto(w, r, &task); err != nil {
		return
	}

	t, err := svc.tasks.Create(r.Context(), task)
	if err != nil {
		web.JsonError(w, http.StatusInternalServerError, err)
		return
	}

	web.Json(w, http.StatusOK, t)
}

func (svc Service) Update(w http.ResponseWriter, r *http.Request) {
	var task gsp.Task

	if err := web.ReadJsonBodyInto(w, r, &task); err != nil {
		return
	}

	err := svc.tasks.Update(r.Context(), task)
	if err != nil {
		web.JsonError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (svc Service) MarkDone(w http.ResponseWriter, r *http.Request) {
	id, err := web.ChiIDGetter(r)
	if err != nil {
		web.JsonError(w, http.StatusBadRequest, err)
		return
	}

	err = svc.tasks.MarkDone(r.Context(), id)
	if err != nil {
		web.JsonError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (svc Service) MarkPending(w http.ResponseWriter, r *http.Request) {
	id, err := web.ChiIDGetter(r)
	if err != nil {
		web.JsonError(w, http.StatusBadRequest, err)
		return
	}

	err = svc.tasks.MarkPending(r.Context(), id)
	if err != nil {
		web.JsonError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (svc Service) List(w http.ResponseWriter, r *http.Request) {
	tasks, err := svc.tasks.List(r.Context())
	if err != nil {
		web.JsonError(w, http.StatusInternalServerError, err)
		return
	}

	web.Json(w, http.StatusOK, tasks)
}

func (svc Service) ListDone(w http.ResponseWriter, r *http.Request) {
	tasks, err := svc.tasks.ListDone(r.Context())
	if err != nil {
		web.JsonError(w, http.StatusInternalServerError, err)
		return
	}

	web.Json(w, http.StatusOK, tasks)
}

func (svc Service) ListPending(w http.ResponseWriter, r *http.Request) {
	tasks, err := svc.tasks.ListPending(r.Context())
	if err != nil {
		web.JsonError(w, http.StatusInternalServerError, err)
		return
	}

	web.Json(w, http.StatusOK, tasks)
}
