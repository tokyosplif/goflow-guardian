package config

import "errors"

const (
	errMsgKafkaBrokersRequired = "kafka brokers are required"
)

type Kafka struct {
	Brokers []string `env:"KAFKA_BROKERS" envSeparator:"," envDefault:"localhost:9092"`
	Topic   string   `env:"KAFKA_TOPIC" envDefault:"security-alerts"`
}

func (c *Kafka) Validate() error {
	if len(c.Brokers) == 0 {
		return errors.New(errMsgKafkaBrokersRequired)
	}
	return nil
}
