package api

import (
	"net/http"
	"time"
)

var HTTPClient = &http.Client{
	Timeout: 100 * time.Second, 
}
