package handlers

import (
	"api/internal/api"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// CreateChatHandler creates a chat using Rails API
func CreateChatHandler(c *gin.Context) {
	appToken := c.Param("application_id")

	chat, err := api.CreateChat(appToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, chat)
}

// GetChatHandler retrieves a chat from Rails API
func GetChatHandler(c *gin.Context) {
	appToken := c.Param("application_id")
	chatNumber, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat ID"})
		return
	}

	chat, err := api.GetChat(appToken, chatNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Chat not found"})
		return
	}

	c.JSON(http.StatusOK, chat)
}
