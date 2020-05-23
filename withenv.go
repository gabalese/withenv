package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]
	dotEnvFileName := args[0]
	programWithArguments := args[1:]

	cmd := exec.Command(strings.Join(programWithArguments, " "))
	envMap, err := godotenv.Read(dotEnvFileName)

	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot read env file %s!", dotEnvFileName))
	}

	var executionEnvironment []string
	executionEnvironment = append(os.Environ())

	for key, value := range envMap {
		environmentVariable := fmt.Sprintf("%s=%s", key, value)
		executionEnvironment = append(executionEnvironment, environmentVariable)
	}

	cmd.Env = executionEnvironment

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		println(err.Error())
	}
}
