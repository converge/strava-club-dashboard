package services

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

type StravaService struct {}

func NewStrava() (StravaService, error) {
	return StravaService{}, nil
}

// SendKafkaMessage a
// produce a message to a specific topic
func SendKafkaMessage () error {

	conf := kafka.ConfigMap{}
	conf["bootstrap.sever"] = "localhost:9092"
	producer, err := kafka.NewProducer(&conf)
	if err != nil {
		log.Println(err)
		return err
	}
	defer producer.Close()

	topic := "strava-activities"
	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic,
			Partition: kafka.PartitionAny,
		},
		Key: []byte("a"),
		Value: []byte("b"),
	}

	err = producer.Produce(&msg, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	producer.Flush(15 * 3000)
	return nil
}
