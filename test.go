package main

import (
	"bufio"
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*
Так же требуется поддерживать функционал fork/exec-команд
*Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).
*/

func kill(s string, out io.Writer) {
	pid, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintf(out, "kill: %v\n", err)
		return
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintf(out, "kill: %v\n", err)
		return
	}

	err = process.Kill()
	if err != nil {
		fmt.Fprintf(out, "kill: %v\n", err)
		return
	}

	fmt.Fprintf(out, "killed proc %d\n", pid)
}

func processList(out io.Writer) {
	processList, err := ps.Processes()
	if err != nil {
		fmt.Fprintf(out, "ps: %v\n", err)
		return
	}
	var b strings.Builder
	for x := range processList {
		var process ps.Process
		process = processList[x]
		b.WriteString(fmt.Sprintf("%d\t%s\n", process.Pid(), process.Executable()))
	}
	fmt.Fprintln(out, b.String())
}

func main() {
	matches, _ := filepath.Glob("/proc/*/exe")
	fmt.Println(matches)
	scanner := bufio.NewScanner(os.Stdin)
	done := false
	var cmd []string
	for scanner.Scan() && !done {
		cmd = strings.Split(scanner.Text(), " ")
		switch cmd[0] {
		case "cd":
			if len(cmd) != 1 {
				err := os.Chdir(cmd[1])
				if err != nil {
					fmt.Fprintf(os.Stdout, "cd: %v\n", err)
				}
			}

		case "pwd":
			path, err := os.Getwd()
			if err != nil {
				fmt.Fprintf(os.Stdout, "pwd: %v\n", err)
			}
			fmt.Fprintln(os.Stdout, path)
		case "echo":
			if len(cmd) > 1 {
				fmt.Fprintln(os.Stdout, cmd[1:])
			}
		case "kill":
			if len(cmd) > 0 {
				kill(cmd[1], os.Stdout)
			}
		case "ps":
			processList(os.Stdout)
		default:
			command := exec.Command(cmd[0], cmd[1:]...)
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
			command.Run()

		}
	}

}
