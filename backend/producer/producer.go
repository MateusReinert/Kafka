package producer

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func SendMessage(broker string, topic string, message string) {
	// Configuração do produtor
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Criando o produtor
	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatal("Erro ao criar produtor Kafka: ", err)
	}
	defer producer.Close()

	// Preparando a mensagem
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// Enviando a mensagem
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatal("Erro ao enviar mensagem: ", err)
	}

	// Exibindo o resultado
	fmt.Printf("Mensagem enviada para a partição %d, offset %d\n", partition, offset)
}
