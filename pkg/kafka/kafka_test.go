package kafka

import (
	"context"
	k "github.com/segmentio/kafka-go"
	"testing"
	"time"
)

// TestKafkaConnection calls new Apache Kafka Producer
func TestKafkaConnection(t *testing.T) {

	uri := "localhost:29092"
	ctx, cf := context.WithTimeout(context.Background(), time.Second*3)

	defer cf()

	producer := newKafkaProducer(HEALTHCHECK_TOPIC, uri)
	defer producer.Close()

	err := producer.WriteMessages(
		ctx,
		k.Message{Value: []byte("OK")},
	)
	if err != nil {
		t.Error(err)
	}
}
