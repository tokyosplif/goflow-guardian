package publisher

import (
	"context"
	"encoding/json"

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
	b, _ := json.Marshal(v)
	return k.writer.WriteMessages(ctx, kafka.Message{Key: []byte(v.Key), Value: b})
}

func (k *Kafka) Ping(ctx context.Context) error {
	conn, err := kafka.DialContext(ctx, "tcp", k.brokers[0])
	if err != nil {
		return err
	}

	return conn.Close()
}
