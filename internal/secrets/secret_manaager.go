package secrets

import (
	"context"
	"fmt"

	sm "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

type SecretManager struct {
	ProjectId string
}

func (s *SecretManager) GetSecret(ctx context.Context, key string) (string, error) {

	c, err := sm.NewClient(ctx)
	if err != nil {
		return "", err
	}

	defer c.Close()

	name := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", s.ProjectId, key)
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}
	result, err := c.AccessSecretVersion(ctx, req)

	if err != nil {
		return "", err
	}

	return string(result.Payload.Data), nil
}
