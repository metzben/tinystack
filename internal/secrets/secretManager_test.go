package secrets

import (
	"os"
	"testing"

	"github.com/metzben/tinystack/internal/config"
	"github.com/metzben/tinystack/pkg/assert"
)

var projectID string

func init() {
	envFile, _ := os.Open("../../.env")
	config, _ := config.Load(envFile)
	projectID = config.GCPProjectID
}

// go test -v '-run=^TestCreateSecretClient$'
func TestCreateSecretClient(t *testing.T) {
	t.Log("Testing... CreateSecretClient")
	_, err := NewGoogleSecretsClient(projectID)
	if err != nil {
		t.Errorf("Error creating google secret client: %v", err)
	}
}

// go test -v '-run=^TestCreateRetrieveDeleteSecret$'
func TestCreateRetrieveDeleteSecret(t *testing.T) {
	t.Log("Testing... CreateSecret")
	t.Logf("projectID: %v", projectID)

	client, err := NewGoogleSecretsClient(projectID)
	if err != nil {
		t.Errorf("Error creating google secret client: %v", err)
	}
	secret := "test"
	secretID := "test-secret"

	versionName, createErr := client.CreateSecret(secretID, []byte(secret))
	if createErr != nil {
		t.Errorf("error creating secret: %v", createErr)
	}
	t.Logf("secretName: %v", versionName)
	// Get Secret
	secretInBytes, secretRetrieveErr := client.GetSecret(versionName)
	if secretRetrieveErr != nil {
		t.Errorf("error retrieving secret: %v", secretRetrieveErr)
	}

	assert.Equal(t, string(secretInBytes), secret)

	deleteErr := client.DeleteSecret(projectID, secretID)
	if deleteErr != nil {
		t.Errorf("error deleting secret: %v", deleteErr)
	}
}
