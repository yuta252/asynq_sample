package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// a list of task types
const (
	TypeEmailDelivery = "email:deliver"
	TypeImageResize   = "image:resize"
)

type EmailDeliveryPayload struct {
	UserID     int
	TemplateID string
}

type ImageResizePayload struct {
	SourceURL string
}

func NewEmailDeliveryTask(userID int, tmplID string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{UserID: userID, TemplateID: tmplID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

func NewImageResizeTask(src string) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageResizePayload{SourceURL: src})
	if err != nil {
		return nil, err
	}
	// enqueue時間をオーバーライドできる
	return asynq.NewTask(TypeEmailDelivery, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
}

// タスクを処理する関数は asynq.HandlerFunc interface を満たす必要がある
// 必ずしも関数である必要はない
func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Marshal failed: %v: %w", p.UserID, asynq.SkipRetry)
	}
	log.Printf("Sending email to User: user_id=%d, template_id=%s", p.UserID, p.TemplateID)
	return nil
}

// asynq.Handler interface を満たす必要がある
type ImageProcessor struct{}

func (processor *ImageProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p ImageResizePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Resizing image: src=%s", p.SourceURL)
	return nil
}

func NewImageProcessor() *ImageProcessor {
	return &ImageProcessor{}
}
