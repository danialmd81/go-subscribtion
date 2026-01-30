package main

import (
	"os"

	"github.com/danialmd81/my-subscribtion/subs"
	"github.com/danialmd81/my-subscribtion/telegram"
)

func main() {
	// Save current proxy settings
	origHTTPProxy := os.Getenv("HTTP_PROXY")
	origHTTPSProxy := os.Getenv("HTTPS_PROXY")

	// Set proxy
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:2080")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:2080")

	// Run telegram with proxy
	telegram.Run()

	// Restore original proxy settings
	os.Setenv("HTTP_PROXY", origHTTPProxy)
	os.Setenv("HTTPS_PROXY", origHTTPSProxy)

	// Continue with other services
	subs.Run()
	// all.Run()

}
