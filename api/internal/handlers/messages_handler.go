package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"api/internal/api"

	"github.com/gin-gonic/gin"
)

var HTTPClient = &http.Client{}


func CreateMessageHandler(c *gin.Context) {
	applicationID := c.Param("application_id")
	chatID, err := strconv.Atoi(c.Param("chat_id")) // Convert chatID to int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chat_id"})
		return
	}

	// Define request struct
	var req struct {
		Body   string `json:"body"`
		Sender string `json:"sender"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert struct to JSON payload
	jsonData, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode JSON"})
		return
	}

	// Make HTTP request
	url := fmt.Sprintf("http://backend:3000/applications/%s/chats/%d/messages", applicationID, chatID)
	resp, err := HTTPClient.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// Read response from Rails API
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Define struct for expected response
	var res struct {
		Status string `json:"status"`
		ChatID int    `json:"chat_id"`
	}

	// Parse JSON response
	if err := json.Unmarshal(body, &res); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	// Return response to client
	c.JSON(resp.StatusCode, res)
}



func SearchMessagesHandler(c *gin.Context) {
	applicationToken := c.Param("application_id")
	chatNumber, err := strconv.Atoi(c.Param("chat_id")) // Convert chatID to int

	// Read request body for search query
	var requestBody struct {
		Query string `json:"query"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call the search function
	messages, err := api.SearchMessages(applicationToken, chatNumber, requestBody.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the search results
	c.JSON(http.StatusOK, messages)
}
