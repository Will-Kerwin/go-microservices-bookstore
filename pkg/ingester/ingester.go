package ingester

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Define a kafka Ingester
type Ingester[T any] struct {
	consumer *kafka.Consumer
	topic    string
}

// create a new ingester
func New[T any](addr string, groupID string, topic string) (*Ingester[T], error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{"bootstrap.servers": addr, "group.id": groupID})

	if err != nil {
		return nil, err
	}
	return &Ingester[T]{consumer, topic}, nil
}

func (i *Ingester[T]) Ingest(ctx context.Context) (<-chan T, error) {

	fmt.Printf("Starting ingestion for %s\n", i.topic)
	if err := i.consumer.SubscribeTopics([]string{i.topic}, nil); err != nil {
		return nil, err
	}

	ch := make(chan T, 1)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				i.consumer.Close()
			default:
			}

			msg, err := i.consumer.ReadMessage(-1)
			if err != nil {
				log.Println("Consumer error:", err)
				continue
			}

			var event T
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				log.Println("Failed to unmarshal event:", err)
				continue
			}

			ch <- event
		}
	}()

	return ch, nil

}
