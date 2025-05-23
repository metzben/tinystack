package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	app.Logger.Info().Msg("home route hit")
	fmt.Fprintln(w, "This is the home route!")
}

func (app *Application) HandleUserName(w http.ResponseWriter, r *http.Request) {
	app.Logger.Info().Msg("user route hit")
	name := r.PathValue("name")
	fmt.Fprintln(w, "yo name is: ", name)
}

type AnthropicMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AnthropicRequest struct {
	Model     string             `json:"model"`
	MaxTokens int                `json:"max_tokens"`
	Messages  []AnthropicMessage `json:"messages"`
}

type ContentItem struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type UsageInfo struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

type AnthropicResponse struct {
	Content      []ContentItem `json:"content"`
	ID           string        `json:"id"`
	Model        string        `json:"model"`
	Role         string        `json:"role"`
	StopReason   string        `json:"stop_reason"`
	StopSequence *string       `json:"stop_sequence"`
	Type         string        `json:"type"`
	Usage        UsageInfo     `json:"usage"`
}

// api endpoint that will accept a user prompt
// /v1/messages/{AnthropicMessage}
func (app *Application) anthropicMessages(w http.ResponseWriter, r *http.Request) {
	// parse incoming json - AnthropicMessage struct
	anthropicMessages := &[]AnthropicMessage{}
	jsonErr := json.NewDecoder(r.Body).Decode(anthropicMessages)
	if jsonErr != nil {
		app.Logger.Err(jsonErr).Msgf("error decoding to struct %v\n", jsonErr)
		return
	}
	anthropicUrl := "https://api.anthropic.com/v1/messages"
	apiKey, secretMgrErr := app.SecretManager.GetSecret(app.Configuration.PathToAnthropicAPIKey)
	if secretMgrErr != nil {
		app.Logger.Err(secretMgrErr).Msg("error calling gcp secret manager")
		return
	}

	anthropicRequest := AnthropicRequest{
		Model:     "claude-3-7-sonnet-20250219",
		MaxTokens: 1024,
		Messages:  *anthropicMessages,
	}

	payload, marshalErr := json.Marshal(anthropicRequest)
	if marshalErr != nil {
		app.Logger.Err(marshalErr).Msg("error parsing struct")
		return
	}

	payloadInBytes := bytes.NewBuffer(payload)
	req, _ := http.NewRequest("POST", anthropicUrl, payloadInBytes)

	// Set required headers
	req.Header.Set("x-api-key", string(apiKey))
	req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("content-type", "application/json")

	// create http client to call anthropic endpoint
	client := &http.Client{Timeout: 10 * time.Second}
	resp, respErr := client.Do(req)
	if respErr != nil {
		app.Logger.Err(respErr).Msg("error executing request")
		return
	}
	defer resp.Body.Close()

	// Parse response to AnthropicResponse struct
	var anthropicResponse AnthropicResponse
	if err := json.NewDecoder(resp.Body).Decode(&anthropicResponse); err != nil {
		app.Logger.Err(err).Msg("error decoding response from Anthropic API")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)

	if err := json.NewEncoder(w).Encode(anthropicResponse); err != nil {
		app.Logger.Err(err).Msg("error encoding response")
	}
}
