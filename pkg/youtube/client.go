package youtube

import (
	"context"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	ytcore "google.golang.org/api/youtube/v3"
)

// NewService creates a new YouTube service from the given credentials
//
// Get these from CredentialsFromBase64
func NewService(ctx context.Context, creds *google.Credentials) (*ytcore.Service, error) {
	clientOpt := option.WithCredentials(creds)
	ytSvc, err := ytcore.NewService(ctx, clientOpt)
	if err != nil {
		return nil, err
	}
	return ytSvc, nil

}
