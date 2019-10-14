package youtube

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
)

// TokenStore is the interface to hold and provide tokens. You can create or
// revoke tokens at
// https://console.developers.google.com/apis/library/youtube.googleapis.com
type TokenStore interface {
	Load() (*oauth2.Token, error)
	Save(*oauth2.Token) error
}

type tokenFileStore struct {
	fileName string
}

func NewTokenFileStore(fileName string) TokenStore {
	return &tokenFileStore{fileName: fileName}
}

func (s *tokenFileStore) Load() (*oauth2.Token, error) {
	f, err := os.Open(s.fileName)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()

	return t, err
}

func (s *tokenFileStore) Save(token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", s.fileName)
	f, err := os.OpenFile(s.fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(token)
}
