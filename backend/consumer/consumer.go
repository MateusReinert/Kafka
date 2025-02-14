package consumer

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func ReceiveMessages(broker string, topic string, sendToClients func(string)) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	client, err := sarama.NewClient([]string{broker}, config)
	if err != nil {
		log.Fatalf("Erro ao criar cliente Kafka: %v\n", err)
	}
	defer client.Close()

	partitions, err := client.Partitions(topic)
	if err != nil {
		log.Fatalf("Erro ao obter partições do tópico %s: %v\n", topic, err)
	}
	fmt.Printf("Partições do tópico %s: %v\n", topic, partitions)

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatalf("Erro ao criar consumidor: %v\n", err)
	}
	defer consumer.Close()

	for _, partition := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("Erro ao consumir da partição %d: %v\n", partition, err)
		}
		defer partitionConsumer.Close()

		go func(partition int32) {
			for message := range partitionConsumer.Messages() {
				fmt.Printf("Mensagem recebida da partição %d: %s\n", partition, string(message.Value))
				sendToClients(string(message.Value)) // Envia a mensagem para o front-end via WebSocket
			}
		}(partition)
	}

	select {}
}
