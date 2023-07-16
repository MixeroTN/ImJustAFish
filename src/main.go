package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filename := "src/file.html"
	path := filepath.Join(wd, filename)

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", "start", "chrome.exe", "--autoplay-policy=no-user-gesture-required", "-kiosk", "file:///"+fmt.Sprint(path), "--disable-features=PreloadMediaEngagementData, MediaEngagementBypassAutoplayPolicies")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	default:
		fmt.Printf("Sorry, this tool is not supported on %s yet.\n", runtime.GOOS)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error launching web browser:", err)
		return
	}
}
