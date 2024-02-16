package task

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
)

// 任务类型
const (
	TypeWelcomeEmail  = "email:welcome"
	TypeReminderEmail = "email:reminder"
)

// 任务执行所需的数据
type emailTaskPayload struct {
	UserID int
}

// NewWelcomeEmailTask 创建一个发送欢迎邮件的任务
func NewWelcomeEmailTask(userID int) (*asynq.Task, error) {
	payload, err := json.Marshal(emailTaskPayload{UserID: userID})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeWelcomeEmail, payload), nil
}

// NewReminderEmailTask 创建一个发送提醒邮件的任务
func NewReminderEmailTask(userID int) (*asynq.Task, error) {
	payload, err := json.Marshal(emailTaskPayload{UserID: userID})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeReminderEmail, payload), nil
}

// Handler 处理任务
// 处理welcome任务
func HandleWelcomeEmailTask(ctx context.Context, t *asynq.Task) error {
	var p emailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	// 发送邮件任务处理逻辑
	log.Printf(" [*] Send Welcome Email to User %d", p.UserID)

	return nil
	//模拟不成功的情况
	//return fmt.Errorf("failed to send welcome email to user %d", p.UserID)
}

// 处理reminder任务
func HandleReminderEmailTask(ctx context.Context, t *asynq.Task) error {
	var p emailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	// 发送邮件任务处理逻辑
	log.Printf(" [*] Send Reminder Email to User %d", p.UserID)
	return nil
}
