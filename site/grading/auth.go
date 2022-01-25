package grading

import (
	ent "170-ag/ent/generated"
	"170-ag/site/roles"
	"context"
	"crypto/hmac"
	"errors"
	"os"
)

var referenceToken = os.Getenv("SCALER_TOKEN")

func withAuthorizedContext(ctx context.Context, client *ent.Client, token string) (context.Context, error) {
	if !hmac.Equal([]byte(referenceToken), []byte(token)) {
		return nil, errors.New("invalid token")
	}
	return roles.WithGradingContext(ctx, client), nil
}
