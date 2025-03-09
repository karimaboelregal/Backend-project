package api

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
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
