package services

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rs/zerolog/log"
)

type StravaService struct {
	*kafka.Producer
	*kafka.Consumer
}

func NewStrava(kafkaProducer *kafka.Producer, kafkaConsumer *kafka.Consumer) StravaService {
	return StravaService{
		kafkaProducer,
		kafkaConsumer,
	}
}

type RiderClubActivity struct {
	ResourceState int `json:"resource_state"`
	Athlete       struct {
		ResourceState int    `json:"resource_state"`
		Firstname     string `json:"firstname"`
		Lastname      string `json:"lastname"`
	} `json:"athlete"`
	Name               string  `json:"name"`
	Distance           float64 `json:"distance"`
	MovingTime         int     `json:"moving_time"`
	ElapsedTime        int     `json:"elapsed_time"`
	TotalElevationGain float64 `json:"total_elevation_gain"`
	Type               string  `json:"type"`
}

// ProduceKafkaMessage a
// produce a message to a specific topic
func (service StravaService) ProduceKafkaMessage(rca RiderClubActivity) error {

	topic := "strava-club-activities"
	riderId := fmt.Sprintf("%s %s", rca.Athlete.Firstname, rca.Athlete.Lastname)
	distance := fmt.Sprintln("distance: %T", rca.Distance)
	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:   []byte(riderId),
		Value: []byte(distance),
	}

	go func() {
		for e := range service.Producer.Events() {
			switch event := e.(type) {
			case *kafka.Message:
				if event.TopicPartition.Error != nil {
					log.Error().Msg("delivered failed...")
				} else {
					log.Info().Msgf("delivered message to %v", event.TopicPartition)
				}
			}
		}
	}()

	err := service.Producer.Produce(&msg, nil)
	if err != nil {
		log.Err(err)
		return err
	}
	service.Producer.Flush(15 * 3000)
	return nil
}

func (service StravaService) GetKafkaMessage() error {

	//log.Info().Msg("loading get kafka...")
	//// config
	//bootstrapServer := os.Getenv("BOOTSTRAP_SERVER")
	//
	//conf := kafka.ConfigMap{}
	//conf["bootstrap.servers"] = bootstrapServer
	//conf["group.id"] = "kafka-go-getting-started"
	//conf["auto.offset.reset"] = "earliest"
	//
	//c, err := kafka.NewConsumer(&conf)
	//if err != nil {
	//	log.Info().Msg("failed to create consumer")
	//	panic(err)
	//}
	//log.Info().Msg("consumer started")

	topic := "strava-club-activities"
	err := service.Consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Err(err)
		return err
	}

	// Process messages
	for {
		log.Info().Msg("waiting new messages...")
		msg, err := service.Consumer.ReadMessage(-1)
		if err == nil {
			log.Info().Str("Message: ", string(msg.Value)).Msgf("Partition: %s", msg.TopicPartition)
		} else {
			log.Error().Msgf("consumer error: %v (%v)", err, msg)
		}
	}
	//service.Close()
	return nil
}
