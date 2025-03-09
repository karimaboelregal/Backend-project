package main

import (
	"api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Application & Chat Routes
	r.POST("/applications", handlers.CreateApplicationHandler)
	r.GET("/applications/:application_id", handlers.GetApplicationHandler)
	r.POST("/applications/:application_id/chats", handlers.CreateChatHandler)
	r.GET("/applications/:application_id/chats/:chat_id", handlers.GetChatHandler)

	// Message Routes
	r.POST("/applications/:application_id/chats/:chat_id/messages", handlers.CreateMessageHandler)
	r.POST("/applications/:application_id/chats/:chat_id/messages/search", handlers.SearchMessagesHandler)

	// Start the server
	r.Run(":8080") // Run on port 8080
}
