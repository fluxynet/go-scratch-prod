package main

import (
	"database/sql"
	"log"
	"net/http"

	gsp "github.com/fluxynet/go-scratch-prod"
	taskRepositoryG "github.com/fluxynet/go-scratch-prod/gorm/tasks"
	taskRepositoryS "github.com/fluxynet/go-scratch-prod/sql/tasks"
	taskService "github.com/fluxynet/go-scratch-prod/task"
	taskHttp "github.com/fluxynet/go-scratch-prod/web/task"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

type Config struct {
	IsDev      bool
	ListenAddr string
	DSN        string
}

func main() {
	router := chi.NewRouter()

	cfg, err := loadConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}

	if cfg.IsDev {
		router.Mount("/api", devRouter(*cfg))
	} else {
		router.Mount("/api", prodRouter(*cfg))
	}

	log.Println("Server started on: " + cfg.ListenAddr)

	http.ListenAndServe(cfg.ListenAddr, router)
}

func loadConfig() (*Config, error) {
	c := Config{
		IsDev:      true,
		ListenAddr: "127.0.0.1:9000",
		DSN:        ":memory:",
	}

	//TODO load config from environment

	return &c, nil
}

func gormTaskRepo(cfg Config) gsp.TaskRepo {
	db, err := gorm.Open(sqlite.Open(cfg.DSN))
	if err != nil {
		log.Fatalln("failed to connect to database: " + err.Error())
	}

	db.AutoMigrate(&gsp.Task{})

	return taskRepositoryG.New(db)
}

func sqliteTaskRepo(cfg Config) gsp.TaskRepo {
	db, err := sql.Open("sqlite", cfg.DSN)
	if err != nil {
		log.Fatalln("failed to connect to database: " + err.Error())
	}

	return taskRepositoryS.New(db)
}

func prodRouter(cfg Config) http.Handler {
	r := chi.NewRouter()

	taskSvc := taskService.New(gormTaskRepo(cfg))
	taskApi := taskHttp.New(taskSvc)

	r.Post("/tasks", taskApi.Create)
	r.Put("/tasks/{id}", taskApi.Update)
	r.Put("/tasks/{id}/done", taskApi.MarkDone)
	r.Put("/tasks/{id}/pending", taskApi.MarkPending)
	r.Get("/tasks", taskApi.List)
	r.Get("/tasks/done", taskApi.ListDone)
	r.Get("/tasks/pending", taskApi.ListPending)

	return r
}

func devRouter(cfg Config) http.Handler {
	p := prodRouter(cfg)

	return p
}
