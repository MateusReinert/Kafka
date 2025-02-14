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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)

func sendMessageToClients(consumerID string, message string) {
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
	broker := "kafka:9092"
	topic := "meu-topico"

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
	}))

	router.GET("/ws", handleWebSocket)

	router.POST("/start-consumer", func(c *gin.Context) {
		consumerID := fmt.Sprintf("consumer-%d", time.Now().UnixNano())

		go consumer.ReceiveMessages(broker, topic, func(msg string) {
			sendMessageToClients(consumerID, msg)
		})

		c.JSON(http.StatusOK, gin.H{"status": "Consumer iniciado", "consumerID": consumerID})
	})

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
