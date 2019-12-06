package youtube

import (
	"context"
	"encoding/base64"

	"golang.org/x/oauth2/google"
)

const youtubeUploadScope = "https://www.googleapis.com/auth/youtube.upload"

// CredentialsFromBase64 fetches google credentials from a base64
// encoded service account JSON file, with YouTube upload scope.
//
// To create a new service account, go here:
//
// https://console.cloud.google.com/apis/credentials?project=go-in-5-minutes&folder&organizationId
//
// and click "Create Credentials" in the top left. Select "Service Account Key"
// and download the JSON file. Then, base64 encode this file. If you're lazy
// like me, you can encode it here: https://www.base64encode.org/
func CredentialsFromBase64(ctx context.Context, bites []byte) (*google.Credentials, error) {
	decoded, err := base64.StdEncoding.DecodeString(string(bites))
	if err != nil {
		return nil, err
	}
	return google.CredentialsFromJSON(ctx, []byte(decoded), youtubeUploadScope)
}
