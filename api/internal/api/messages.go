package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Message struct
type Message struct {
	Number int    `json:"message_number"`
	Body   string `json:"body"`
}

// CreateMessage sends a request to Rails API to create a message
func CreateMessage(applicationToken string, chatNumber int, body string, sender string) (map[string]string, error) {
	url := fmt.Sprintf("http://backend:3000/applications/%s/chats/%d/messages", applicationToken, chatNumber)

	requestBody, err := json.Marshal(map[string]string{
		"body":   body,
		"sender": sender,
	})
	if err != nil {
		return nil, err
	}

	resp, err := HTTPClient.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response map[string]string
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// SearchMessages sends a POST request to Rails with a query
func SearchMessages(applicationToken string, chatNumber int, query string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("http://backend:3000/applications/%s/chats/%d/messages/search", applicationToken, chatNumber)

	// Create JSON body with the query
	requestBody, err := json.Marshal(map[string]string{"query": query})
	if err != nil {
		return nil, err
	}

	// Send POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var messages []map[string]interface{}
	if err := json.Unmarshal(body, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
