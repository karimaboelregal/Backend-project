package api

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"net/http"
)

type Application struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

func CreateApplication(name string) (*Application, error) {
	url := "http://backend:3000/applications"
	reqBody := fmt.Sprintf(`{"name": "%s"}`, name)
	resp, err := HTTPClient.Post(url, "application/json", strings.NewReader(reqBody))
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
	var app Application
	if err := json.Unmarshal(body, &app); err != nil {
		return nil, err
	}

	return &app, nil
}



type ApplicationGetApi struct {
	ID         int    `json:"id"`
	Token      string `json:"token"`
	Name       string `json:"name"`
	ChatsCount int    `json:"chats_count"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func GetApplication(applicationToken string) (*ApplicationGetApi, error) {
	url := fmt.Sprintf("http://backend:3000/applications/%s", applicationToken)

	resp, err := HTTPClient.Get(url) 
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch application: status %d", resp.StatusCode)
	}


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}


	var app ApplicationGetApi
	if err := json.Unmarshal(body, &app); err != nil {
		return nil, err
	}

	return &app, nil
}

