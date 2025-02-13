package main

import (
	"fmt"
	"kafka/consumer"
	"kafka/producer"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Definir um upgrader para WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool) // Armazenar clientes conectados

func sendMessageToClients(consumerID string, message string) {
	// Corrigindo o uso de fmt.Sprintf
	jsonMsg := fmt.Sprintf(`{"consumerID": "%s", "message": "%s"}`, consumerID, message)
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(jsonMsg))
		if err != nil {
			log.Printf("Erro ao enviar mensagem para cliente: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Erro ao upgrade para WebSocket:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro no WebSocket"})
		return
	}
	defer conn.Close()

	clients[conn] = true
	defer delete(clients, conn)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Erro ao ler mensagem do cliente: %v", err)
			break
		}
	}
}

func main() {
	// Definindo o broker e o tópico Kafka
	broker := "kafka:9092" // Kafka está no container chamado "kafka"
	topic := "meu-topico"

	// Inicializando o Gin e as rotas
	router := gin.Default()

	// Adicionando middleware CORS para permitir acesso do localhost:3001
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,                                                // Permitir qualquer origem
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}, // Permitir métodos HTTP padrão
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
	}))

	// Rota para WebSocket
	router.GET("/ws", handleWebSocket)

	// Rota para iniciar o consumo de mensagens
	router.POST("/start-consumer", func(c *gin.Context) {
		consumerID := fmt.Sprintf("consumer-%d", time.Now().UnixNano()) // Gera um ID único

		go consumer.ReceiveMessages(broker, topic, func(msg string) {
			sendMessageToClients(consumerID, msg)
		})

		c.JSON(http.StatusOK, gin.H{"status": "Consumer iniciado", "consumerID": consumerID})
	})

	// Rota para enviar mensagem para o Kafka
	router.POST("/send-message", func(c *gin.Context) {
		var msg struct {
			Message string `json:"message"`
		}
		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mensagem inválida"})
			return
		}
		go producer.SendMessage(broker, topic, msg.Message)
		c.JSON(http.StatusOK, gin.H{"status": "Mensagem enviada para o Kafka"})
	})

	log.Println("Servidor HTTPS iniciado na porta 4000...")
	log.Fatal(router.Run(":4000"))

}
