package exec

import (
	"context"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestNewExec(t *testing.T) {
	e := NewExec()
	_ = e
}

func TestExec_LookPath(t *testing.T) {
	e := NewExec()

	// Look for a command that should exist on all platforms
	var testCmd string
	if runtime.GOOS == "windows" {
		testCmd = "cmd"
	} else {
		testCmd = "sh"
	}

	path, err := e.LookPath(testCmd)
	if err != nil {
		t.Errorf("LookPath(%q) returned error: %v", testCmd, err)
	}
	if path == "" {
		t.Errorf("LookPath(%q) returned empty path", testCmd)
	}
}

func TestExec_LookPath_NotFound(t *testing.T) {
	e := NewExec()

	_, err := e.LookPath("nonexistent-command-that-should-not-exist")
	if err == nil {
		t.Error("LookPath for nonexistent command should return error")
	}
}

func TestExec_NewCommand(t *testing.T) {
	e := NewExec()

	// Use go version as a safe command that exists
	cmd := e.NewCommand("go", WithArgs("version"))
	if cmd == nil {
		t.Fatal("NewCommand returned nil")
	}
}

func TestCmd_Run(t *testing.T) {
	e := NewExec()

	cmd := e.NewCommand("go", WithArgs("version"))
	err := cmd.Run()
	if err != nil {
		t.Errorf("Run() returned error: %v", err)
	}
}

func TestCmd_Output(t *testing.T) {
	e := NewExec()

	cmd := e.NewCommand("go", WithArgs("version"))
	output, err := cmd.Output()
	if err != nil {
		t.Errorf("Output() returned error: %v", err)
	}
	if len(output) == 0 {
		t.Error("Output() returned empty output")
	}
	if !strings.Contains(string(output), "go version") {
		t.Errorf("Output() = %q, want to contain 'go version'", string(output))
	}
}

func TestCmd_CombinedOutput(t *testing.T) {
	e := NewExec()

	cmd := e.NewCommand("go", WithArgs("version"))
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("CombinedOutput() returned error: %v", err)
	}
	if len(output) == 0 {
		t.Error("CombinedOutput() returned empty output")
	}
}

func TestCmd_Start_Wait(t *testing.T) {
	e := NewExec()

	cmd := e.NewCommand("go", WithArgs("version"))

	err := cmd.Start()
	if err != nil {
		t.Errorf("Start() returned error: %v", err)
	}

	err = cmd.Wait()
	if err != nil {
		t.Errorf("Wait() returned error: %v", err)
	}
}

func TestCmd_Env(t *testing.T) {
	e := NewExec()

	cmd := e.NewCommand("go", WithArgs("version"), WithEnv("TEST_VAR", "test_value"))

	env := cmd.Env()
	if len(env) == 0 {
		t.Error("Env() returned empty slice")
	}

	found := false
	for _, e := range env {
		if e == "TEST_VAR=test_value" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Env() didn't contain expected environment variable")
	}
}

func TestCmd_Dir(t *testing.T) {
	e := NewExec()

	tmpDir := t.TempDir()
	cmd := e.NewCommand("go", WithArgs("version"), WithDir(tmpDir))

	dir := cmd.Dir()
	if dir != tmpDir {
		t.Errorf("Dir() = %q, want %q", dir, tmpDir)
	}
}

func TestCmd_Args(t *testing.T) {
	e := NewExec()

	cmd := e.NewCommand("go", WithArgs("version"))
	args := cmd.Args()
	if len(args) < 2 {
		t.Errorf("Args() returned %d args, want at least 2", len(args))
	}
	if args[1] != "version" {
		t.Errorf("Args()[1] = %q, want 'version'", args[1])
	}
}

func TestCmd_Path(t *testing.T) {
	e := NewExec()

	cmd := e.NewCommand("go", WithArgs("version"))
	path := cmd.Path()
	if path == "" {
		t.Error("Path() returned empty string")
	}
}

func TestCmd_String(t *testing.T) {
	e := NewExec()

	cmd := e.NewCommand("go", WithArgs("version"))
	str := cmd.String()
	if str == "" {
		t.Error("String() returned empty string")
	}
}

func TestCmd_Environ(t *testing.T) {
	e := NewExec()

	cmd := e.NewCommand("go", WithArgs("version"), WithEnv("TEST", "value"))
	environ := cmd.Environ()
	if len(environ) == 0 {
		t.Error("Environ() returned empty slice")
	}
}

func TestCmd_StdoutPipe(t *testing.T) {
	e := NewExec()

	cmd := e.NewCommand("go", WithArgs("version"))
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		t.Errorf("StdoutPipe() returned error: %v", err)
	}
	if pipe == nil {
		t.Error("StdoutPipe() returned nil")
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		t.Fatalf("Start() returned error: %v", err)
	}

	// Read from pipe
	buf := make([]byte, 1024)
	n, _ := pipe.Read(buf)
	if n == 0 {
		t.Error("StdoutPipe didn't receive any data")
	}

	pipe.Close()
	cmd.Wait()
}

func TestCmd_StderrPipe(t *testing.T) {
	e := NewExec()

	// Run a command that outputs to stderr
	cmd := e.NewCommand("go", WithArgs("help", "nonexistent"))
	pipe, err := cmd.StderrPipe()
	if err != nil {
		t.Errorf("StderrPipe() returned error: %v", err)
	}
	if pipe == nil {
		t.Error("StderrPipe() returned nil")
	}

	cmd.Start()
	pipe.Close()
	cmd.Wait()
}

func TestCmd_StdinPipe(t *testing.T) {
	e := NewExec()

	// Use a command that can read from stdin
	var testCmd string
	var testArgs []string
	if runtime.GOOS == "windows" {
		testCmd = "cmd"
		testArgs = []string{"/C", "echo test"}
	} else {
		testCmd = "cat"
		testArgs = []string{}
	}

	cmd := e.NewCommand(testCmd, WithArgs(testArgs...))
	pipe, err := cmd.StdinPipe()
	if err != nil {
		t.Errorf("StdinPipe() returned error: %v", err)
	}
	if pipe == nil {
		t.Error("StdinPipe() returned nil")
	}

	cmd.Start()
	pipe.Close()
	cmd.Wait()
}

func TestCmd_Process(t *testing.T) {
	e := NewExec()

	cmd := e.NewCommand("go", WithArgs("version"))
	if err := cmd.Start(); err != nil {
		t.Fatalf("Start() returned error: %v", err)
	}

	proc := cmd.Process()
	if proc == nil {
		t.Error("Process() returned nil after Start()")
	}
	if proc.Pid == 0 {
		t.Error("Process().Pid = 0, want > 0")
	}

	cmd.Wait()
}

func TestCmd_ProcessState(t *testing.T) {
	e := NewExec()

	cmd := e.NewCommand("go", WithArgs("version"))
	if err := cmd.Run(); err != nil {
		t.Fatalf("Run() returned error: %v", err)
	}

	state := cmd.ProcessState()
	if state == nil {
		t.Error("ProcessState() returned nil after Run()")
	}
	if !state.Exited() {
		t.Error("ProcessState().Exited() = false, want true")
	}
}

func TestCmd_Stdin_Stdout_Stderr(t *testing.T) {
	e := NewExec()

	var stdin strings.Reader
	var stdout, stderr strings.Builder

	cmd := e.NewCommand("go", WithArgs("version"),
		WithStdin(&stdin),
		WithStdout(&stdout),
		WithStderr(&stderr))

	if cmd.Stdin() != &stdin {
		t.Error("Stdin() didn't return set value")
	}
	if cmd.Stdout() != &stdout {
		t.Error("Stdout() didn't return set value")
	}
	if cmd.Stderr() != &stderr {
		t.Error("Stderr() didn't return set value")
	}
}

func TestWithContext(t *testing.T) {
	e := NewExec()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := e.NewCommand("go", WithArgs("version"), WithContext(ctx))
	if err := cmd.Run(); err != nil {
		t.Errorf("Run() with context returned error: %v", err)
	}
}

func TestWithContext_Timeout(t *testing.T) {
	e := NewExec()

	// Create a context that expires very quickly
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()

	// Give context time to expire
	time.Sleep(10 * time.Millisecond)

	// Use a command that would normally take a while
	var testCmd string
	var testArgs []string
	if runtime.GOOS == "windows" {
		testCmd = "ping"
		testArgs = []string{"127.0.0.1", "-n", "10"}
	} else {
		testCmd = "sleep"
		testArgs = []string{"10"}
	}

	cmd := e.NewCommand(testCmd, WithArgs(testArgs...), WithContext(ctx))
	err := cmd.Run()
	if err == nil {
		t.Error("Expected error from expired context, got nil")
	}
}

func TestWithExtraOptions(t *testing.T) {
	e := NewExec()

	// Test WithWaitDelay
	cmd := e.NewCommand("go", WithArgs("version"), WithWaitDelay(1*time.Second))
	if err := cmd.Run(); err != nil {
		t.Errorf("Run() with WaitDelay returned error: %v", err)
	}
}

func TestErrorTypes(t *testing.T) {
	// Test that Error and ExitError types are accessible
	var _ *Error
	var _ *ExitError
}
