package main

import (
	"github.com/bitfield/script"
)

func main() {
	execCmdWithMatchWithReject("kubectl get pod", "workflow", "bundle").Stdout()
}

func execCmd(command string, match string, reject string) *script.Pipe {
	return script.Exec(command).Column(1)
}

func execCmdWithMatch(command string, match string) *script.Pipe {
	return script.Exec(command).Match(match).Column(1)
}

func execCmdWithMatchWithReject(command string, match string, reject string) *script.Pipe {
	return script.Exec(command).Match(match).Reject(reject).Column(1)
}
