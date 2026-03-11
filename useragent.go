package gocardless

import (
	"fmt"
	"runtime"
)

const (
	// client library version
	ClientLibVersion = "6.1.0"
)

var userAgent string

func initUserAgent() {
	goVersion := runtime.Version()
	userAgent = fmt.Sprintf("gocardless-pro-go/%s go/%s", ClientLibVersion, goVersion)
}
