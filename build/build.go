package build

import (
	"fmt"
	"runtime"
)

// Version the application version
var Version = "development"

// SHA is the git reference used to build this package
var SHA = "unknown"

// BuildDate is when it was build
var BuildDate = "unknown"

// License is the official Open Source Initiave license abbreciation
var License = "Apache-2.0"

// Set to default to no TLS
var TLS = "true"

// Set to disable signing and certificates in the protocol
var Secure = "true"

func String() string {
	return fmt.Sprintf("version: %s\n\nlicense: %s\nbuilt: %s\nsha: %s\ntls: %s\nsecure: %s\ngo: %s", Version, License, BuildDate, SHA, TLS, Secure, runtime.Version())
}