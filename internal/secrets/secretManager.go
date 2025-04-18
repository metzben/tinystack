package secrets

import (
	"context"
	"fmt"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

type SecretManager interface {
	CreateSecret(secretName string, secretValue []byte) (string, error)
	GetSecret(versionName string) ([]byte, error)
	DeleteSecret(projectID, secretID string) error
	Close() error
}

// The secretmanager client will use your default application credentials.
// Clients should be reused instead of created as needed.
// The methods of Client are safe for concurrent use by multiple goroutines.
// The returned client must be Closed when it is done being used.

type GoogleSecretsClient struct {
	ProjectID string
	Client    *secretmanager.Client
}

func NewGoogleSecretsClient(projectID string) (GoogleSecretsClient, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return GoogleSecretsClient{}, fmt.Errorf("failed to setup client: %v", err)
	}

	return GoogleSecretsClient{
		ProjectID: projectID,
		Client:    client,
	}, nil
}

func (gsc GoogleSecretsClient) CreateSecret(secretName string, secretValue []byte) (string, error) {
	// Create the request to add the secret
	secretReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", gsc.ProjectID),
		SecretId: secretName,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}
	// Set context with timeout
	ctxWithTimout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	secret, secretCreateErr := gsc.Client.CreateSecret(ctxWithTimout, secretReq)
	if secretCreateErr != nil {
		return "", fmt.Errorf("failed to create secret: %v", secretCreateErr)
	}
	// Create the request to add the secret version
	versionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: secretValue,
		},
	}
	// Set context with timeout
	ctxwto, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Call the API
	version, addSecretVersionErr := gsc.Client.AddSecretVersion(ctxwto, versionReq)
	if addSecretVersionErr != nil {
		return "", fmt.Errorf("failed to add secret version: %v", addSecretVersionErr)
	}
	// Return the version name
	return version.Name, nil
}

func (smc GoogleSecretsClient) GetSecret(versionName string) ([]byte, error) {
	// Build the request
	getSecretRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: versionName,
	}
	// Set context with timeout
	ctxWithTimout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Call the API
	result, err := smc.Client.AccessSecretVersion(ctxWithTimout, getSecretRequest)
	if err != nil {
		return nil, fmt.Errorf("error retrieving secret %v", err)
	}
	return result.Payload.Data, nil
}

func (smc GoogleSecretsClient) DeleteSecret(projectID, secretID string) error {
	secretName := fmt.Sprintf("projects/%s/secrets/%s", projectID, secretID)

	tctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Build the request for deleting the secret version
	req := &secretmanagerpb.DeleteSecretRequest{
		Name: secretName,
	}
	// Call the API to delete the secret.
	if err := smc.Client.DeleteSecret(tctx, req); err != nil {
		return fmt.Errorf("failed to delete secret: %v", err)
	}
	return nil
}

func (smc GoogleSecretsClient) Close() error {
	return smc.Client.Close()
}
