package main

import (
	"os/exec"
)

func main() {
	anotherPackagePath := "../anotherDirectory"
	cmd := exec.Command("./goRun.sh", anotherPackagePath)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	cmd.Wait()
}
