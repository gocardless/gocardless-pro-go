package gocardless

import (
	"fmt"
	"runtime"
)

const (
	// client library version
	clientLibVersion = "4.5.0"
)

var userAgent string

func initUserAgent() {
	goVersion := runtime.Version()
	userAgent = fmt.Sprintf("gocardless-pro-go/%s go/%s", clientLibVersion, goVersion)
}
