package gocardless

import (
	"fmt"
	"runtime"
)

const (
	// client library version
	clientLibVersion = "5.2.0"
)

var userAgent string

func initUserAgent() {
	goVersion := runtime.Version()
	userAgent = fmt.Sprintf("gocardless-pro-go/%s go/%s", clientLibVersion, goVersion)
}
