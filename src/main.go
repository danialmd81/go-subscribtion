package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/danialmd81/my-subscribtion/all"
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
	all.Run()

	// Git commit and push
	if err := gitCommitAndPush(); err != nil {
		fmt.Printf("Git error: %v\n", err)
	}
}

func gitCommitAndPush() error {
	commands := [][]string{
		{"git", "add", "."},
		{"git", "commit", "-m", "Update proxy configurations"},
		{"git", "push", "origin", "main"},
	}

	for _, args := range commands {
		cmd := exec.Command(args[0], args[1:]...)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("command failed: %v", err)
		}
		fmt.Printf("[INFO] Executed: %v\n", args)
	}

	return nil
}
