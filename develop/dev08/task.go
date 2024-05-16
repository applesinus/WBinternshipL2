package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	Ex8()
}

// Ex8 - test task 8
func Ex8() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "\\quit" {
			break
		}

		// split input into commands
		commands := strings.Split(input, "|")

		// execute each command in the pipeline
		executePipeline(commands)
	}
}

func executePipeline(commands []string) {
	var prevOutput io.Reader
	var err error

	for _, cmdStr := range commands {
		args := strings.Fields(strings.TrimSpace(cmdStr))

		switch args[0] {
		case "cd", "pwd", "echo", "kill":
			executeCommand(args)
		default:
			cmd := exec.Command(args[0], args[1:]...)

			if prevOutput != nil {
				cmd.Stdin = prevOutput
			}

			var output []byte
			output, err = cmd.Output()

			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка при выполнении команды:", err)
				return
			}

			fmt.Print(string(output))

			prevOutput = bytes.NewReader(output)
		}
	}
}

func executeCommand(args []string) {
	// if there are no arguments, do nothing
	if len(args) == 0 {
		return
	}

	command := args[0]

	switch command {
	case "cd":
		// change current directory
		if len(args) < 2 {
			fmt.Println("Использование: cd <директория>")
			return
		}
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка при смене директории:", err)
		}
	case "pwd":
		// show current directory
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка при получении текущей директории:", err)
			return
		}
		fmt.Println(cwd)
	case "echo":
		// show arguments
		fmt.Println(strings.Join(args[1:], " "))
	case "kill":
		// kill process
		if len(args) < 2 {
			fmt.Println("Использование: kill <pid>")
			return
		}
		pid := args[1]
		cmd := exec.Command("kill", pid)
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка при убивании процесса:", err)
		}
	default:
		// any other command
		cmd := exec.Command(command, args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка при выполнении команды:", err)
		}
	}
}
