package shell

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/mitchellh/go-ps"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

//Shell - структура параметров работы
type Shell struct {
	out             io.Writer
	in              io.Reader
	nameUser        string
	nameSystem      string
	colorSystemUser string
	colorDir        string
	colorInput      string
	buffPipe        *bytes.Buffer
	isPipe          bool
}

// NewShell Создать Объект Shell
func NewShell(out io.Writer, in io.Reader) *Shell {
	return &Shell{out: out, in: in}
}

// SetColors задаем цвета вывода в Shell
func (s *Shell) SetColors(system, pathDir, input string) {
	s.colorSystemUser = system
	s.colorDir = pathDir
	s.colorInput = input
}

// SetSystem задаем имя системы и пользователя для отображения в Shell
func (s *Shell) SetSystem(sys, usr string) {
	s.nameSystem = sys
	s.nameUser = usr
}

// Run Запустить shell
func (s *Shell) Run() {
	if err := s.ScanLines(); err != nil {
		if _, err := fmt.Fprintln(s.out, err); err != nil {
			log.Fatalln(err)
		}
	}
}

// PrefixComm вывод перед строкой ввода директории в которой находимся и пользователя
func (s *Shell) PrefixComm() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(s.out, "%v%v@%v:%v%v%v$ ", s.colorSystemUser, s.nameUser, s.nameSystem, s.colorDir, dir, s.colorInput)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

// ScanLines читаем команды до \quit
func (s *Shell) ScanLines() error {
	sc := bufio.NewScanner(s.in)
	if err := s.PrefixComm(); err != nil {
		return err
	}
	for sc.Scan() && sc.Text() != "\\quit" {
		line := sc.Text()
		if err := s.CheckFork(line); err != nil {
			return err
		}
		if err := s.PrefixComm(); err != nil {
			return err
		}
	}

	if sc.Err() != nil {
		os.Exit(1)
	}
	return nil
}

// CheckPipes Обработка строки на наличие конвеера
func (s *Shell) CheckPipes(line string) error {
	lineCmd := strings.Split(line, "|")
	if len(lineCmd) > 1 { // если конвеер
		s.buffPipe = new(bytes.Buffer) //буфер между пайпами
		s.isPipe = true
		for index, comm := range lineCmd {
			if index != 0 { // если не первая команда то разбиваем ее ( для args)
				commSl := strings.Fields(comm)

				if len(commSl) > 1 { // если args есть у команды, то заменяем их на значения из буфера
					commSlNew := make([]string, 2, 2)
					commSlNew[0], commSlNew[1] = commSl[0], s.buffPipe.String()
					commSl = commSlNew
				} else {
					commSl = append(commSl, s.buffPipe.String()) // просто добавление args из буфера

				}
				comm = strings.Join(commSl, " ")
			}
			s.buffPipe.Reset()
			if index == len(lineCmd)-1 { // если команда последняя - вывод в stdout, а не в буффер
				s.isPipe = false
			}

			if err := s.caseCommand(comm); err != nil {

				if _, err = fmt.Fprintln(s.out, err); err != nil {
					return err
				}
			}

		}

	} else { // если нет конвеера
		if err := s.caseCommand(line); err != nil {
			if _, err = fmt.Fprintln(s.out, err); err != nil {
				return err
			}
		}
	}
	return nil
}

// CheckFork Обработка строки на наличие fork()
func (s *Shell) CheckFork(line string) error {
	line = strings.TrimRight(line, " ")
	if strings.Contains(line, "&") {
		if id, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0); err != 0 {
			os.Exit(1)
		} else if id == 0 {
			line = strings.TrimSuffix(line, "&")
			if _, err := fmt.Fprintf(s.out, "[%v]\n", os.Getpid()); err != nil {
				return err
			}
			if err := s.CheckPipes(line); err != nil {
				return err
			}
			if _, err := fmt.Fprintf(s.out, "[%v] Завершен\n", os.Getpid()); err != nil {
				return err
			}
			os.Exit(0)
		}

	} else {
		if err := s.CheckPipes(line); err != nil {
			return err
		}
	}
	return nil
}

func (s *Shell) cd(param string) error {
	if err := os.Chdir(param); err != nil {
		return err
	}
	return nil
}

func (s *Shell) pwd() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	if s.isPipe {
		if _, err := fmt.Fprintln(s.buffPipe, dir); err != nil {
			return err
		}
	} else {
		if _, err := fmt.Fprintln(s.out, dir); err != nil {
			return err
		}
	}
	return nil
}

func (s *Shell) kill(procIds []string) []error {
	var errs []error
	for _, pid := range procIds { // может быть много PID у команды
		if id, err := strconv.Atoi(pid); err != nil {
			err = fmt.Errorf("%v pid: %v", err, id)
		} else {
			if errKill := syscall.Kill(id, syscall.SIGTERM); errKill != nil { // по умолчанию в Шелл сигнал SIGTERM
				errKill = fmt.Errorf("%v pid: %v", errKill, id)
				errs = append(errs, errKill)
			}
		}
	}
	return errs
}

func (s *Shell) echo(args []string, fullComm string) error {
	strStart := args[0]
	strEnd := args[len(args)-1]
	if strStart[0] == '"' && strEnd[len(strEnd)-1] == '"' { // поддержка вывода фиксированной строки
		start, end := strings.Index(fullComm, "\""), strings.LastIndex(fullComm, "\"")
		var newArgs []string
		newArgs = append(newArgs, fullComm[start+1:end])
		args = newArgs
	}
	output := s.out
	if s.isPipe {
		output = s.buffPipe
	}
	if _, err := fmt.Fprintln(output, strings.Join(args, " ")); err != nil {
		return err
	}
	return nil
}

func (s *Shell) ps() error {
	procs, err := ps.Processes()
	if err != nil {
		return err
	}
	output := s.out
	if s.isPipe {
		output = s.buffPipe
	}
	if _, err := fmt.Fprintf(output, "%50v%10v%10v\n", "CMD", "PID", "PPI"); err != nil {
		return err
	}
	for _, proc := range procs {
		if _, err := fmt.Fprintf(output, "%50v%10v%10v\n", proc.Executable(), proc.Pid(), proc.PPid()); err != nil {
			return err
		}
	}
	return nil
}

// выполнение внешних команд
func (s *Shell) exec(args []string) error {
	var command *exec.Cmd
	if len(args) == 1 {
		command = exec.Command(args[0])
	} else {
		command = exec.Command(args[0], args[1:]...)
	}
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if s.isPipe {
		command.Stdout = s.buffPipe
	}
	if err := command.Run(); err != nil {
		return err
	}
	return nil
}

func (s *Shell) caseCommand(comm string) error {
	if setComm := strings.Fields(comm); len(setComm) != 0 {
		switch setComm[0] {
		case "cd":
			if len(setComm) == 2 {
				if err := s.cd(setComm[1]); err != nil {
					if _, err := fmt.Fprintln(s.out, err); err != nil {
						return err
					}
				}
			} else {
				return errChangeDir
			}
		case "pwd":
			if len(setComm) == 1 {
				if err := s.pwd(); err != nil {
					return err
				}
			} else {
				return errPWD
			}
		case "kill":
			if len(setComm) >= 2 {
				errs := s.kill(setComm[1:])
				for _, err := range errs {
					if _, err := fmt.Fprintf(s.out, "%v\n", err); err != nil {
						return err
					}
				}
			} else {
				return errKill
			}
		case "echo":
			if len(setComm) >= 2 {
				if err := s.echo(setComm[1:], comm); err != nil {
					return err
				}
			} else {
				return errEcho
			}
		case "ps":
			if len(setComm) == 1 {
				if err := s.ps(); err != nil {
					return err
				}
			} else {
				return errPS
			}
		case "exec":
			if len(setComm) >= 2 {
				if err := s.exec(setComm[1:]); err != nil {
					return err
				}
			} else {
				return errExec
			}
		}
	}
	return nil
}
