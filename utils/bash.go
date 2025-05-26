package utils

import (
    "bytes"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

// RunCommand executes an inline Bash command string and returns combined stdout/stderr output.
func RunCommand(command string) (string, error) {
    cmd := exec.Command("bash", "-c", command)
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &out

    err := cmd.Run()
    return strings.TrimSpace(out.String()), err
}

// RunScript executes a Bash script file located at scriptPath with optional args.
// Returns combined stdout/stderr output.
func RunScript(scriptPath string, args ...string) (string, error) {
    cmd := exec.Command("bash", append([]string{scriptPath}, args...)...)
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &out

    err := cmd.Run()
    return strings.TrimSpace(out.String()), err
}

// RunCommandWithEnv runs a Bash command with additional environment variables.
func RunCommandWithEnv(command string, envVars map[string]string) (string, error) {
    cmd := exec.Command("bash", "-c", command)
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &out

    env := os.Environ()
    for k, v := range envVars {
        env = append(env, fmt.Sprintf("%s=%s", k, v))
    }
    cmd.Env = env

    err := cmd.Run()
    return strings.TrimSpace(out.String()), err
}

// RunAndCapture splits output and error separately.
func RunAndCapture(command string) (stdout string, stderr string, err error) {
    cmd := exec.Command("bash", "-c", command)
    var outBuf, errBuf bytes.Buffer
    cmd.Stdout = &outBuf
    cmd.Stderr = &errBuf

    err = cmd.Run()
    return strings.TrimSpace(outBuf.String()), strings.TrimSpace(errBuf.String()), err
}
