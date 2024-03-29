package main

import (
	"bytes"
	"os/exec"
	"strconv"
)

type PowerShell struct {
	powerShell string
}

func New() *PowerShell {
	ps, _ := exec.LookPath("powershell")
	return &PowerShell{
		powerShell: ps,
	}
}

func (p *PowerShell) runCmd(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}

func execPowerShell(podName string) (stdOut string, stdErr string, err error) {
	posh := New()

	kubectl, _ := exec.LookPath("kubectl")
	arguments := "exec -it " + podName + " -- sh"

	processCmds := `
	$newProcess = new-object System.Diagnostics.ProcessStartInfo "PowerShell";
	$newProcess.Arguments = "` + arguments + `";
	$newProcess.FileName = "` + kubectl + `";
	$newProcess.CreateNoWindow = "true";
	$process = [System.Diagnostics.Process]::Start($newProcess);
	exit	
	`

	return posh.runCmd(processCmds)
}

func execPowerShellContext(context string) (stdOut string, stdErr string, err error) {
	posh := New()

	kubectl, _ := exec.LookPath("kubectl")
	arguments := "config use-context " + context

	processCmds := `
	$newProcess = new-object System.Diagnostics.ProcessStartInfo "PowerShell";
	$newProcess.Arguments = "` + arguments + `";
	$newProcess.FileName = "` + kubectl + `";
	$newProcess.CreateNoWindow = "false";
	$process = [System.Diagnostics.Process]::Start($newProcess);
	exit	
	`

	return posh.runCmd(processCmds)
}

func execPowerShellDelete(podName string) (stdOut string, stdErr string, err error) {
	posh := New()

	kubectl, _ := exec.LookPath("kubectl")
	arguments := "delete pod " + podName

	processCmds := `
	$newProcess = new-object System.Diagnostics.ProcessStartInfo "PowerShell";
	$newProcess.Arguments = "` + arguments + `";
	$newProcess.FileName = "` + kubectl + `";
	$newProcess.CreateNoWindow = "false";
	$process = [System.Diagnostics.Process]::Start($newProcess);
	exit	
	`

	return posh.runCmd(processCmds)
}

func execPowerShellCopyFiles(podName string, pathLog string, pathPod string) (stdOut string, stdErr string, err error) {
	posh := New()
	kubectl, _ := exec.LookPath("kubectl")
	arguments := "cp " + podName + ":" + pathPod + " " + pathLog

	processCmds := `
	$newProcess = new-object System.Diagnostics.ProcessStartInfo "PowerShell";
	$newProcess.Arguments = "` + arguments + `";
	$newProcess.FileName = "` + kubectl + `";
	$newProcess.CreateNoWindow = "false";
	$process = [System.Diagnostics.Process]::Start($newProcess);
	exit	
	`

	return posh.runCmd(processCmds)
}

func execPowerShellZipLogFolder(podName string, pathPod string) (stdOut string, stdErr string, err error) {
	posh := New()

	arguments := "exec -it " + podName + " -- bash -c "
	tarArgs := strconv.Quote("tar -cf " + PATH_LOG_POD + " " + pathPod)
	arguments = "kubectl " + arguments + tarArgs

	return posh.runCmd(arguments)
}
