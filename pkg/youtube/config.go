package youtube

import (
	"context"

	ytcore "google.golang.org/api/youtube/v3"
)

// Config holds common configuration values for YouTube interaction
type Config struct {
	credsBase64 []byte
	channelID   string
}

func NewConfig(credsBase64 []byte, chanID string) Config {
	return Config{
		credsBase64: credsBase64,
		channelID:   chanID,
	}
}

func (c Config) Service(ctx context.Context) (*ytcore.Service, error) {
	googleCreds, err := CredentialsFromBase64(
		ctx,
		c.credsBase64,
	)
	if err != nil {
		return nil, err
	}
	return NewService(ctx, googleCreds)

}

func (c Config) ChannelID() string {
	return c.channelID
}
