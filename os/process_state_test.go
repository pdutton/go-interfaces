package os

import (
	"os"
	"os/exec"
	"testing"
)

func TestWrapProcessState(t *testing.T) {
	// Run a simple command to get a ProcessState
	cmd := exec.Command("true")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	realPS := cmd.ProcessState
	if realPS == nil {
		t.Fatal("ProcessState is nil")
	}

	wrapped := WrapProcessState(realPS)

	// Verify it implements the interface
	var _ ProcessState = wrapped
}

func TestProcessState_Nub(t *testing.T) {
	cmd := exec.Command("true")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	realPS := cmd.ProcessState
	wrapped := WrapProcessState(realPS)

	if wrapped.Nub() != realPS {
		t.Error("Nub() did not return the original ProcessState")
	}
}

func TestProcessState_SuccessfulCommand(t *testing.T) {
	cmd := exec.Command("true")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	ps := WrapProcessState(cmd.ProcessState)

	if ps.ExitCode() != 0 {
		t.Errorf("ExitCode() = %d, want 0", ps.ExitCode())
	}

	if !ps.Exited() {
		t.Error("Exited() = false, want true")
	}

	if ps.Pid() <= 0 {
		t.Errorf("Pid() = %d, want > 0", ps.Pid())
	}

	if !ps.Success() {
		t.Error("Success() = false, want true")
	}

	str := ps.String()
	if str == "" {
		t.Error("String() returned empty string")
	}
}

func TestProcessState_FailedCommand(t *testing.T) {
	cmd := exec.Command("false")
	err := cmd.Run()
	if err == nil {
		t.Skip("'false' command did not fail as expected")
	}

	ps := WrapProcessState(cmd.ProcessState)

	if ps.ExitCode() == 0 {
		t.Error("ExitCode() = 0, want non-zero for failed command")
	}

	if !ps.Exited() {
		t.Error("Exited() = false, want true")
	}

	if ps.Success() {
		t.Error("Success() = true, want false for failed command")
	}
}

func TestProcessState_Sys(t *testing.T) {
	cmd := exec.Command("true")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	ps := WrapProcessState(cmd.ProcessState)

	// Sys() returns platform-specific data, just verify it doesn't panic
	sys := ps.Sys()
	if sys == nil {
		t.Log("Sys() returned nil (may be expected on some platforms)")
	}
}

func TestProcessState_SysUsage(t *testing.T) {
	cmd := exec.Command("true")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	ps := WrapProcessState(cmd.ProcessState)

	// SysUsage() returns platform-specific data, just verify it doesn't panic
	usage := ps.SysUsage()
	if usage == nil {
		t.Log("SysUsage() returned nil (may be expected on some platforms)")
	}
}

func TestProcessState_Times(t *testing.T) {
	cmd := exec.Command("true")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	ps := WrapProcessState(cmd.ProcessState)

	// Times should be non-negative
	userTime := ps.UserTime()
	if userTime < 0 {
		t.Errorf("UserTime() = %v, want >= 0", userTime)
	}

	systemTime := ps.SystemTime()
	if systemTime < 0 {
		t.Errorf("SystemTime() = %v, want >= 0", systemTime)
	}
}

func TestProcessState_WithProcessWait(t *testing.T) {
	// Test integration with Process.Wait()
	osf := NewOS()

	cmd := exec.Command("true")
	err := cmd.Start()
	if err != nil {
		t.Fatalf("Command Start() failed: %v", err)
	}

	// Wrap the os.Process and call Wait()
	proc := WrapProcess(cmd.Process)
	ps, err := proc.Wait()
	if err != nil {
		t.Fatalf("Wait() error = %v", err)
	}

	if ps == nil {
		t.Fatal("Wait() returned nil ProcessState")
	}

	if !ps.Success() {
		t.Error("Wait() ProcessState.Success() = false, want true")
	}

	// Verify we can access the underlying ProcessState
	nub := ps.Nub()
	if nub == nil {
		t.Error("Nub() returned nil")
	}

	// Silence unused variable warning
	_ = osf
}

func TestProcessState_FindProcessAndWait(t *testing.T) {
	// Start a command that we can find and wait on
	cmd := exec.Command("sleep", "0.1")
	err := cmd.Start()
	if err != nil {
		t.Fatalf("Command Start() failed: %v", err)
	}

	pid := cmd.Process.Pid

	osf := NewOS()
	proc, err := osf.FindProcess(pid)
	if err != nil {
		t.Fatalf("FindProcess() error = %v", err)
	}

	ps, err := proc.Wait()
	if err != nil {
		t.Fatalf("Wait() error = %v", err)
	}

	if ps.Pid() != pid {
		t.Errorf("ProcessState.Pid() = %d, want %d", ps.Pid(), pid)
	}
}

func TestProcessState_StartProcessAndWait(t *testing.T) {
	osf := NewOS()

	// Find the 'true' command path
	truePath, err := exec.LookPath("true")
	if err != nil {
		t.Skipf("'true' command not found: %v", err)
	}

	proc, err := osf.StartProcess(truePath, []string{"true"}, &os.ProcAttr{})
	if err != nil {
		t.Fatalf("StartProcess() error = %v", err)
	}

	ps, err := proc.Wait()
	if err != nil {
		t.Fatalf("Wait() error = %v", err)
	}

	if !ps.Success() {
		t.Error("ProcessState.Success() = false, want true")
	}

	if ps.ExitCode() != 0 {
		t.Errorf("ProcessState.ExitCode() = %d, want 0", ps.ExitCode())
	}
}
