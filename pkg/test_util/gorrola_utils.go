package test_util

import (
	"fmt"
	"os/exec"
)

func RunGorrola() {
	output, err := exec.Command("./gorrola", "run", "--backends", "http://localhost:8080,http://localhost:8081,http://localhost:8082").Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Command Successfully Executed")
	fmt.Println("Result: ", string(output))
}
