package consumer

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

// Função para receber mensagens do Kafka
func ReceiveMessages(broker string, topic string, sendToClients func(string)) {
	// Configuração do consumidor
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	// Conectando ao broker Kafka
	client, err := sarama.NewClient([]string{broker}, config)
	if err != nil {
		log.Fatalf("Erro ao criar cliente Kafka: %v\n", err)
	}
	defer client.Close()

	// Verificando as partições do tópico
	partitions, err := client.Partitions(topic)
	if err != nil {
		log.Fatalf("Erro ao obter partições do tópico %s: %v\n", topic, err)
	}
	fmt.Printf("Partições do tópico %s: %v\n", topic, partitions)

	// Criando o consumidor
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatalf("Erro ao criar consumidor: %v\n", err)
	}
	defer consumer.Close()

	// Consumindo mensagens das partições
	for _, partition := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("Erro ao consumir da partição %d: %v\n", partition, err)
		}
		defer partitionConsumer.Close()

		// Consumindo mensagens
		go func(partition int32) {
			for message := range partitionConsumer.Messages() {
				fmt.Printf("Mensagem recebida da partição %d: %s\n", partition, string(message.Value))
				sendToClients(string(message.Value)) // Envia a mensagem para o front-end via WebSocket
			}
		}(partition)
	}

	// Bloquear o programa principal para continuar ouvindo
	select {}
}
