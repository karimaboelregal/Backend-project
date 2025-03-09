package api

import (
	"encoding/json"
	"fmt"
	"io"
)

// Chat represents a chat object from the Rails API
type Chat struct {
	Number int `json:"chat_number"`
}


// CreateChat makes a request to the Rails API to create a chat
func CreateChat(applicationToken string) (*Chat, error) {
	url := fmt.Sprintf("http://backend:3000/applications/%s/chats", applicationToken)
	resp, err := HTTPClient.Post(url, "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var chat Chat
	if err := json.Unmarshal(body, &chat); err != nil {
		return nil, err
	}

	return &chat, nil
}

type GetChatStructure struct {
	ID            int       `json:"id"`
	ApplicationID int       `json:"application_id"`
	Number        int       `json:"number"`
	MessagesCount *int      `json:"messages_count"` // Use *int to handle null values
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
}

// GetChat fetches a chat by ID
func GetChat(applicationToken string, chatNumber int) (*GetChatStructure, error) {
	url := fmt.Sprintf("http://backend:3000/applications/%s/chats/%d", applicationToken, chatNumber)
	resp, err := HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var chat GetChatStructure
	if err := json.Unmarshal(body, &chat); err != nil {
		return nil, err
	}

	return &chat, nil
}
