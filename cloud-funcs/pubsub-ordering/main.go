package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

func init() {
	functions.CloudEvent("HelloPubSub", helloPubSub)
}

// MessagePublishedData contains the full Pub/Sub message
// See the documentation for more details:
// https://cloud.google.com/eventarc/docs/cloudevents#pubsub
type MessagePublishedData struct {
	Message PubSubMessage
}

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// helloPubSub consumes a CloudEvent message and extracts the Pub/Sub message.
func helloPubSub(ctx context.Context, e event.Event) error {
	var msg MessagePublishedData
	if err := e.DataAs(&msg); err != nil {
		fmt.Println("event.DataAs: ", err)
		return nil
	}

	user := new(User)
	if err := json.Unmarshal(msg.Message.Data, user); err != nil {
		log.Println("Error unmarshalling user:", err, string(msg.Message.Data))
		return nil
	}

	// Set your Google Cloud Platform project ID.
	projectID := os.Getenv("PID")

	// Creates a client.
	client, err := datastore.NewClientWithDatabase(ctx, projectID, os.Getenv("DB"))
	if err != nil {
		log.Println("Failed to create client: %v", err)
		return nil
	}

	defer client.Close()

	// Sets the kind for the new entity.
	kind := os.Getenv("KIND")
	// Sets the name/ID for the new entity.
	name := user.UserId
	// Creates a Key instance.
	taskKey := datastore.NameKey(kind, name, nil)
	// Saves the new entity.
	if _, err := client.Put(ctx, taskKey, user); err != nil {
		log.Println("Failed to save task: %v", err)
		return nil
	}

	fmt.Printf("Saved %v: %v\n", taskKey, user.UserId)
	return nil
}

type User struct {
	UserId    string `json:"user_id,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

func main() {
}
