package main

import (
	"bytes"
	"os/exec"

	"github.com/bitfield/script"
)

func execute(commands ...string) ([]string, error) {

	var argCmd, argMatch, argReject string
	for i, cmd := range commands {
		if i == 0 {
			argCmd = cmd
		}
		if i == 1 {
			argMatch = cmd
		}
		if i == 2 {
			argReject = cmd
		}
	}

	if argMatch != "" {
		return execCmdWithMatchReturnSlice(argCmd, argMatch)
	} else if argMatch != "" && argReject != "" {
		return execCmdWithMatchWithRejectReturnSlice(argCmd, argMatch, argReject)
	}

	return execCmdReturnSlice(argCmd)
}

func execCmd(command string) *script.Pipe {
	return script.Exec(command)
}

func execCmdWithMatch(command string, match string) *script.Pipe {
	return script.Exec(command).Match(match)
}

func execCmdWithMatchWithReject(command string, match string, reject string) *script.Pipe {
	return script.Exec(command).Match(match).Reject(reject)
}

func execCmdReturnSlice(command string) ([]string, error) {
	return execCmd(command).Slice()
}

func execCmdWithMatchReturnSlice(command string, match string) ([]string, error) {
	return execCmdWithMatch(command, match).Slice()
}

func execCmdWithMatchWithRejectReturnSlice(command string, match string, reject string) ([]string, error) {
	return execCmdWithMatchWithReject(command, match, reject).Slice()
}

func execCmcReturnSliceAndColumn(command string, columnNumber int) ([]string, error) {
	return execCmd(command).Column(columnNumber).Slice()
}

func executeKubectlCore(commands ...string) (stdOut string, stdErr string) {
	kubectl, _ := exec.LookPath("kubectl")

	cmd := exec.Command(kubectl, commands...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Run()
	stdOut = stdout.String()
	stdErr = stderr.String()

	return
}
