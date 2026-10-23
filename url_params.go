package gocardless

import (
	"fmt"
	"net/url"
	"strings"
)

// escapeURLParam escapes a value before it is interpolated into a request path. A URL
// parameter is a single path segment, so values that could move the request to a
// different endpoint - path separators, control characters, "." and ".." (escaping can't
// make these safe, since a resolver strips them regardless), and empty values - are
// rejected instead.
func escapeURLParam(key, value string) (string, error) {
	if value == "" {
		return "", fmt.Errorf("no value provided for URL parameter %q", key)
	}

	if value == "." || value == ".." {
		return "", fmt.Errorf(
			"invalid value for URL parameter %q: %q would change which endpoint the request is sent to",
			key, value)
	}

	if strings.IndexFunc(value, isForbiddenURLParamRune) >= 0 {
		return "", fmt.Errorf(
			"invalid value for URL parameter %q: %q contains a character that is not allowed in a path segment",
			key, value)
	}

	return url.PathEscape(value), nil
}

func isForbiddenURLParamRune(r rune) bool {
	return r == '/' || r == '?' || r == '#' || r < 0x20 || r == 0x7f
}
