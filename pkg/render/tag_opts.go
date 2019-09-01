package render

import (
	"fmt"
	"strings"
)

type TagOpts map[string]interface{}

func EmptyOpts() TagOpts {
	return TagOpts{new(map[string]interface{})}
}

func (t *TagOpts) String() string {
	if t == nil {
		return ""
	}
	attrs := []string{}
	for k, v := range *t {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, k, v))
	}
	return strings.Join(attrs, " ")
}

// to2 overwrites to1
func MergeTagOpts(to1, to2 TagOpts) TagOpts {
	ret := to1
	for k, v := range to2 {
		ret[k] = v
	}
	return ret
}
