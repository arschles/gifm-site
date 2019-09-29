package security

import (
	"encoding/base64"
	"fmt"

	"github.com/gobuffalo/buffalo"
)

// TokenFieldName is the field name of the authenticity token in the meta and
// hidden form fields
const TokenFieldName = "authenticity_token"

// AuthenticityToken holds an authenticity token and lets you
// print it out, etc...
type AuthenticityToken struct {
	fmt.Stringer
	rawBytes []byte
}

func (a AuthenticityToken) String() string {
	return base64.StdEncoding.EncodeToString(a.rawBytes)
}

// AuthenticityTokenFromCtx gets the authenticity_token value from the
// session and returns it as a string. An empty string will be returned
// if there is no authenticity_token available
func AuthenticityTokenFromCtx(c buffalo.Context) AuthenticityToken {
	bites := c.Session().Get(TokenFieldName).([]byte)
	return AuthenticityToken{rawBytes: bites}
}
