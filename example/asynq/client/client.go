package main

import (
	"github.com/hibiken/asynq"
	"github.com/nico612/go-project/example/asynq/task"
	"log"
	"time"
)

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	t1, err := task.NewWelcomeEmailTask(42)
	if err != nil {
		log.Fatal(err)
	}
	t2, err := task.NewReminderEmailTask(42)
	if err != nil {
		log.Fatal(err)
	}
	// process the task immediately
	info, err := client.Enqueue(t1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[*] Succeffuly enqueued task: %+v", info)

	// process the task 24 hours later
	info, err = client.Enqueue(t2, asynq.ProcessAt(time.Now().Add(24*time.Hour)))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[*] Succeffuly enqueued task: %+v", info)
}
