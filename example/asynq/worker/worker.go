package main

import (
	"github.com/hibiken/asynq"
	"github.com/nico612/go-project/example/asynq/task"
	"log"
)

func main() {
	srv := asynq.NewServer(asynq.RedisClientOpt{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}, asynq.Config{
		Concurrency: 10,
	})

	// 像路由一样，mux将类型映射到处理程序
	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeWelcomeEmail, task.HandleWelcomeEmailTask)
	mux.HandleFunc(task.TypeReminderEmail, task.HandleReminderEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
