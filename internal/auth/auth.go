package auth

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

func NewAuthClient(at string) *http.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: at},
	)

	return oauth2.NewClient(ctx, ts)
}
