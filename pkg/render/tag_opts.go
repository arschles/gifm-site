package render

import (
	"fmt"
	"strings"
)

type TagOpts map[string]interface{}

func EmptyOpts() TagOpts {
	return TagOpts(map[string]interface{}{})
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

// MergeTagOpts merges two sets of TagOpts together with the below
// rules.
//
// to2 overwrites to1 unless a key is in mergeKeys.
// in that case, the val in to1 is appended to to2 after a space
func MergeTagOpts(to1, to2 TagOpts, mergeKeys ...string) TagOpts {
	if to1 == nil && to2 == nil {
		return TagOpts{}
	} else if to1 == nil {
		return to2
	} else if to2 == nil {
		return to1
	}
	mergeKeySet := map[string]struct{}{}
	for _, mergeKey := range mergeKeys {
		mergeKeySet[mergeKey] = struct{}{}
	}

	ret := to1
	for k, v := range to2 {
		_, isMergeKey := mergeKeySet[k]
		to1Val, to1Exists := to1[k]
		if to1Exists && isMergeKey {
			v = fmt.Sprintf("%s %s", v, to1Val)
		}
		ret[k] = v
	}
	return ret
}
