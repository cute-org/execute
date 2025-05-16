package dataflow

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"execute/internal"
)

type TaskEvent struct {
	TaskID    int `json:"task_id"`
	UserID    int `json:"user_id"`
	Timestamp int `json:"timestamp"`
}

var (
	pubsubClient *pubsub.Client
	topic        *pubsub.Topic
)

func InitPS() {
	projectID := os.Getenv("GCP_PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC_NAME")

	if projectID == "" || topicName == "" {
		log.Fatalf("Missing environment variables: GCP_PROJECT_ID or PUBSUB_TOPIC_NAME")
	}

	err := initPubSub(projectID, topicName)
	if err != nil {
		log.Fatalf("PubSub init failed: %v", err)
	}
}

func initPubSub(projectID, topicID string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("failed to create pubsub client: %w", err)
	}
	pubsubClient = client
	topic = client.Topic(topicID)
	return nil
}

func InsertTaskEvent(taskID, userID int) error {
	_, err := internal.DB.Exec(
		`INSERT INTO task_events (task_id, user_id)
		 VALUES ($1, $2)`,
		taskID, userID,
	)
	if err != nil {
		return fmt.Errorf("failed to insert task event: %w", err)
	}

	// Publish to Pub/Sub
	if pubsubClient != nil && topic != nil {
		event := TaskEvent{
			TaskID:    taskID,
			UserID:    userID,
			Timestamp: int(time.Now().UTC().Unix()),
		}
		payload, err := json.Marshal(event)
		if err != nil {
			log.Printf("failed to marshal event for pubsub: %v", err)
			return nil
		}

		ctx := context.Background()
		result := topic.Publish(ctx, &pubsub.Message{Data: payload})
		if _, err := result.Get(ctx); err != nil {
			log.Printf("failed to publish task event to pubsub: %v", err)
		}
	}

	return nil
}
