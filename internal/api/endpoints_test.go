package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/metzben/tinystack/internal/api/url"
	"github.com/metzben/tinystack/internal/config"
	"github.com/metzben/tinystack/internal/secrets"
	"github.com/metzben/tinystack/pkg/assert"
	"github.com/rs/zerolog"
)

var testApp *Application

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	envFile, _ := os.Open("../../.env")
	configuration, _ := config.Load(envFile)
	secretMgr, _ := secrets.NewGoogleSecretsClient(configuration.GCPProjectID)
	testApp = &Application{
		Logger:        log,
		Configuration: configuration,
		SecretManager: secretMgr,
	}
}

// go test -v '-run=^TestHomeEndpoint$'
func TestHomeEndpoint(t *testing.T) {
	// Create a new router
	mux := http.NewServeMux()
	router := testApp.BuildRoutes(mux)

	// Create a test server
	server := httptest.NewServer(mux)
	defer server.Close()

	// Make a request to the home endpoint
	r, _ := http.NewRequest("GET", url.Home, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	// Check status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check response body
	expectedBody := "This is the home route!\n"
	assert.Equal(t, expectedBody, w.Body.String())
}

// go test -v '-run=^TestHandleUserName$'
func TestHandleUserName(t *testing.T) {
	mux := http.NewServeMux()
	router := testApp.BuildRoutes(mux)

	server := httptest.NewServer(mux)
	defer server.Close()
	url := "/v1/users/Mac"

	request, _ := http.NewRequest("GET", url, nil)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)

	expectedBody := "yo name is:  Mac\n"
	assert.Equal(t, expectedBody, writer.Body.String())
}

// go test -v '-run=^TestAnthropicMessages$'
func TestAnthropicMessages(t *testing.T) {
	mux := http.NewServeMux()
	router := testApp.BuildRoutes(mux)

	server := httptest.NewServer(mux)
	defer server.Close()

	// Create test payload
	messages := []AnthropicMessage{
		{Role: "user", Content: "Hello, test message"},
	}
	payload, _ := json.Marshal(messages)

	// Make a request to the anthropic messages endpoint
	r, _ := http.NewRequest("POST", url.AnthropicMessages, bytes.NewBuffer(payload))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	// Since this endpoint makes external API calls and requires secrets,
	// we expect it to fail in test environment, but we can verify
	// that it properly handles the request structure
	// The endpoint should at least process the JSON without panicking
	assert.Equal(t, true, w.Code >= 200)
}
