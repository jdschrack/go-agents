package secrets

import (
	"context"

	sm "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"

	"github.com/jdschrack/go-agents/internal/log"
)

func GetSecret(ctx context.Context, path string) (string, error) {
	c, err := sm.NewClient(ctx)
	if err != nil {
		return "", err
	}

	defer func() {
		if err := c.Close(); err != nil {
			logger := log.FromContext(ctx)
			logger.Err(err).Msg("failed to close secret manager client")
		}
	}()

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: path,
	}
	result, err := c.AccessSecretVersion(ctx, req)

	if err != nil {
		return "", err
	}

	return string(result.Payload.Data), nil
}
