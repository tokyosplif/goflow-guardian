package publisher

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/tokyosplif/goflow-guardian/internal/config"
	"github.com/tokyosplif/goflow-guardian/internal/domain"
)

type Kafka struct {
	writer  *kafka.Writer
	brokers []string
}

func NewKafka(cfg config.Kafka) *Kafka {
	return &Kafka{
		brokers: cfg.Brokers,
		writer: &kafka.Writer{
			Addr:  kafka.TCP(cfg.Brokers...),
			Topic: cfg.Topic,
		},
	}
}

func (k *Kafka) PublishViolation(ctx context.Context, v domain.Violation) error {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal violation: %w", err)
	}

	err = k.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(v.Key),
		Value: b,
	})
	if err != nil {
		return fmt.Errorf("failed to write message to kafka: %w", err)
	}

	return nil
}

func (k *Kafka) Ping(ctx context.Context) error {
	conn, err := kafka.DialContext(ctx, "tcp", k.brokers[0])
	if err != nil {
		return err
	}

	return conn.Close()
}

func (k *Kafka) Close() error {
	return k.writer.Close()
}
