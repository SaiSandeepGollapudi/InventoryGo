package queue

import (
	"fmt"

	"github.com/IBM/sarama"
)

func ProduceKafkaMessage(topic, message string, producer sarama.SyncProducer) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("error producing message to kafak: %v", err)
	}
	fmt.Printf("Message in stored in partition %d offset %d\n", partition, offset)
	return nil
}
