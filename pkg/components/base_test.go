package components
import (
	"testing"
	"bytes"
)
func TestBase(t *testing.T) {
	base := Base(emptyElt{})
	var b bytes.Buffer
	if err := base.Render(&b, emptyData()); err != nil {
		t.Fatal(err)
	}
	t.Logf(string(b.Bytes()))
}