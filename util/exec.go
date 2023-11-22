package util

import "os/exec"

func IsExistCommand(c string) bool {
	_, err := exec.LookPath(c)
	return err == nil
}

func Execute(c string, args ...string) error {
	cmd := exec.Command(c, args...)
	return cmd.Run()
}
