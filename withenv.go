package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]
	dotEnvFileName := args[0]
	programWithArguments := args[1:]

	cmd := exec.Command(programWithArguments[0], programWithArguments[1:]...)

	file, err := os.Open(dotEnvFileName)

	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot read env file %s!", dotEnvFileName))
	}

	var executionEnvironment []string
	executionEnvironment = append(os.Environ())

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		environmentVariable := scanner.Text()
		executionEnvironment = append(executionEnvironment, environmentVariable)
	}

	_ = file.Close()

	cmd.Env = executionEnvironment

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		println(err.Error())
	}
}
