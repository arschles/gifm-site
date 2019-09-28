package forms

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
)

// TokenFieldName is the field name of the authenticity token in the meta and
// hidden form fields
const TokenFieldName = "authenticity_token"

// AuthenticityToken gets the authenticity_token value from the session
// and returns it as a string. An empty string will be returned
// if there is no authenticity_token available
func AuthenticityToken(c buffalo.Context) string {
	bites := c.Session().Get(TokenFieldName).([]byte)
	return fmt.Sprint(bites)
	// strSlice := make([]string, len(bites))
	// for i, bite := range bites {
	// 	strSlice[i] = string(bite)
	// }
	// return strings.Join(strSlice, "")
}
