package tailscale

/*
Copyright © Telnyx LLC

Package tailscale implements interface to interact with Tailscale CLI
*/

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func GetTailscaleUser() string {
	tailscalePath := tailscalePath()
	result, err := exec.Command(tailscalePath, "status").Output()
	if err != nil {
		log.Fatalf("Failed to run Tailscale CLI: %v", err)
	}

	return parseUsername(result)
}

// this parses the command output and returns username@telnyx.com
// example line
// 100.77.71.41    hostname           marcin@      macOS   -
func parseUsername(result []byte) string {
	first_line := strings.Split(string(result), "\n")[0]
	return strings.Fields(first_line)[2] + "telnyx.com"
}

func tailscalePath() string {
	switch runtime.GOOS {
	case "darwin":
		return "/Applications/Tailscale.app/Contents/MacOS/Tailscale"
	case "linux":
		path, err := exec.LookPath("tailscale")
		if err != nil {
			log.Fatal("Tailscale not found in $PATH")
		}
		return path
	default:
		log.Fatalf("Unsupported OS %v", runtime.GOOS)
		return ""
	}
}
