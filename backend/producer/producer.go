package producer

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func SendMessage(broker string, topic string, message string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatal("Erro ao criar produtor Kafka: ", err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatal("Erro ao enviar mensagem: ", err)
	}

	fmt.Printf("Mensagem enviada para a partição %d, offset %d\n", partition, offset)
}
