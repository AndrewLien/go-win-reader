package gopowershell

import (
	"bytes"
	"os/exec"
)

// Create interface for context
type Context interface {
	Enumerate() (string, string)
}

type PowerShell struct {
	LocalExecPath string
	PowerShellStr string
}

type PowerShellError struct {
	ErrorMessage    string
	PSCtx   PowerShell
}

func (pshell *PowerShell) New() (*PowerShell, *PowerShellError) {
	ps, err := exec.LookPath(pshell.LocalExecPath)
	if err != nil {
		return &PowerShell{pshell.LocalExecPath, pshell.PowerShellStr}, &PowerShellError{err.Error(), *pshell}
	}
	return &PowerShell{
		LocalExecPath: pshell.LocalExecPath,
		PowerShellStr: ps,
	}, nil
}

func (pshell *PowerShell) Execute(args ...string) (stdOut string, shellError *PowerShellError) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(pshell.PowerShellStr, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "Cannot run cmd", &PowerShellError{err.Error(), *pshell}
	}
	stdOut, stdErr := stdout.String(), stderr.String()

	return stdOut, &PowerShellError{stdErr, *pshell}
}



func (PowerShell PowerShell) Enumerate() (string, string) {
	return PowerShell.LocalExecPath,  PowerShell.PowerShellStr
}