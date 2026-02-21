package main

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var version = "1.2.0"

// ================= EMBED =================

//go:embed flutter-tool.sh
var scriptContent []byte

//go:embed internal/busybox
var busyboxBinary []byte

// ================= MAIN =================

func main() {
	if err := run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.ExitCode())
		}

		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {

	if isVersionRequest() {
		printVersion()
		return nil
	}

	if runtime.GOOS == "windows" {
		return errors.New("windows is not supported")
	}

	tmpDir, err := os.MkdirTemp("", "flutter-tool-")
	if err != nil {
		return fmt.Errorf("create temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	scriptPath, err := writeScript(tmpDir)
	if err != nil {
		return err
	}

	busyboxBinDir, err := extractBusybox(tmpDir)
	if err != nil {
		return err
	}

	cmd := buildCommand(scriptPath, busyboxBinDir)

	return cmd.Run()
}

// ================= VERSION =================

func isVersionRequest() bool {
	return len(os.Args) > 1 &&
		(os.Args[1] == "--version" || os.Args[1] == "-v")
}

func printVersion() {
	fmt.Printf(
		"flutter-tool (community-distro) v%s (%s/%s)\n",
		version,
		runtime.GOOS,
		runtime.GOARCH,
	)
}

// ================= SCRIPT =================

func writeScript(tmpDir string) (string, error) {
	path := filepath.Join(tmpDir, "flutter-tool.sh")

	if err := os.WriteFile(path, scriptContent, 0755); err != nil {
		return "", fmt.Errorf("write script: %w", err)
	}

	return path, nil
}

// ================= BUSYBOX =================

func extractBusybox(tmpDir string) (string, error) {

	binDir := filepath.Join(tmpDir, "bin")

	if err := os.MkdirAll(binDir, 0755); err != nil {
		return "", fmt.Errorf("create busybox bin dir: %w", err)
	}

	busyboxPath := filepath.Join(binDir, "busybox")

	// Write busybox binary
	if err := os.WriteFile(busyboxPath, busyboxBinary, 0755); err != nil {
		return "", fmt.Errorf("write busybox binary: %w", err)
	}

	// Install symlinks for applets
	cmd := exec.Command(busyboxPath, "--install", "-s", binDir)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("busybox --install failed: %w", err)
	}

	return binDir, nil
}

// ================= COMMAND =================

func buildCommand(scriptPath, busyboxBinDir string) *exec.Cmd {

	cmd := exec.Command(
		"/bin/bash",
		append([]string{scriptPath}, os.Args[1:]...)...,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmd.Env = buildEnvWithPrependedPath(busyboxBinDir)

	return cmd
}

func buildEnvWithPrependedPath(prepend string) []string {

	env := os.Environ()
	newEnv := make([]string, 0, len(env))

	var pathHandled bool

	for _, e := range env {
		if strings.HasPrefix(e, "PATH=") {
			oldPath := strings.TrimPrefix(e, "PATH=")
			newEnv = append(newEnv, "PATH="+prepend+":"+oldPath)
			pathHandled = true
		} else {
			newEnv = append(newEnv, e)
		}
	}

	if !pathHandled {
		newEnv = append(newEnv, "PATH="+prepend)
	}

	return newEnv
}
