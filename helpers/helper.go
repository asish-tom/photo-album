package helpers

import (
	"encoding/json"
	"github.com/Shopify/sarama"
)

func PublishToKafka(topic string, data interface{}) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	// brokers := []string{"192.168.59.103:9092"}
	brokers := []string{"kafka:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {

		// Should not reach here
		panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			// Should not reach here
			panic(err)
		}
	}()
	d, err := json.Marshal(data)
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(d),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}

}
