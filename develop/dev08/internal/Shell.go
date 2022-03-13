package shell

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/mitchellh/go-ps"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

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

func NewShell(out io.Writer, in io.Reader) *Shell {
	return &Shell{out: out, in: in}
}

func (s *Shell) SetColors(system, pathDir, input string) {
	s.colorSystemUser = system
	s.colorDir = pathDir
	s.colorInput = input
}

func (s *Shell) SetSystem(sys, usr string) {
	s.nameSystem = sys
	s.nameUser = usr
}

func (s *Shell) Run() {
	s.ScanLines()
}
func (s *Shell) PrefixComm() error {
	if dir, err := os.Getwd(); err != nil {
		return err
	} else {
		fmt.Fprintf(s.out, "%v%v@%v:%v%v%v$ ", s.colorSystemUser, s.nameUser, s.nameSystem, s.colorDir, dir, s.colorInput)
	}
	return nil
}

func (s *Shell) ScanLines() {
	sc := bufio.NewScanner(s.in)
	s.PrefixComm()
	for sc.Scan() && sc.Text() != "\\quit" {
		line := sc.Text()
		s.CheckFork(line)
		s.PrefixComm()
	}

	if sc.Err() != nil {
		os.Exit(1)
	}
}
func (s *Shell) CheckPipes(line string) {

	lineCmd := strings.Split(line, "|")
	if len(lineCmd) > 1 {
		s.buffPipe = new(bytes.Buffer)
		s.isPipe = true
		for index, comm := range lineCmd {
			if index != 0 {
				commSl := strings.Fields(comm)
				if len(commSl) > 1 {
					commSlNew := make([]string, 2, 2)
					commSlNew[0], commSlNew[1] = commSl[0], s.buffPipe.String()
					commSl = commSlNew
				} else {
					commSl = append(commSl, s.buffPipe.String())
				}
				comm = strings.Join(commSl, " ")
			}
			if index == len(lineCmd)-1 {
				s.isPipe = false
				s.buffPipe.Reset()
			}
			if err := s.CaseCommand(comm); err != nil {

				fmt.Fprintln(s.out, err)
			}

		}

	} else {
		if err := s.CaseCommand(line); err != nil {
			fmt.Fprintln(s.out, err)
		}
	}
}

func (s *Shell) CheckFork(line string) {
	line = strings.TrimRight(line, " ")
	if strings.Contains(line, "&") {
		if id, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0); err != 0 {
			os.Exit(1)
		} else if id == 0 {
			line = strings.TrimSuffix(line, "&")
			fmt.Fprintf(s.out, "[%v]\n", os.Getpid())
			s.CheckPipes(line)
			fmt.Fprintf(s.out, "[%v] Завершен\n", os.Getpid())
			os.Exit(0)
		}

	} else {
		s.CheckPipes(line)
	}
}

func (s *Shell) cd(param string) error {
	if err := os.Chdir(param); err != nil {
		return err
	}
	return nil
}

func (s *Shell) pwd() error {
	if dir, err := os.Getwd(); err != nil {
		return err
	} else {
		if s.isPipe {
			fmt.Fprintln(s.buffPipe, dir)
		} else {
			fmt.Fprintln(s.out, dir)
		}

	}
	return nil
}

func (s *Shell) kill(procIds []string) []error {
	var errs []error
	for _, pid := range procIds {
		if id, err := strconv.Atoi(pid); err != nil {
			err = fmt.Errorf("%v pid: %v", err, id)
		} else {
			if errKill := syscall.Kill(id, syscall.SIGTERM); errKill != nil {
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
	if strStart[0] == '"' && strEnd[len(strEnd)-1] == '"' {
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
	if procs, err := ps.Processes(); err != nil {
		return err
	} else {
		output := s.out
		if s.isPipe {
			output = s.buffPipe
		}
		fmt.Fprintf(output, "%50v%10v%10v\n", "CMD", "PID", "PPI")
		for _, proc := range procs {
			fmt.Fprintf(output, "%50v%10v%10v\n", proc.Executable(), proc.Pid(), proc.PPid())
		}
	}
	return nil
}

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

func (s *Shell) CaseCommand(comm string) error {
	if setComm := strings.Fields(comm); len(setComm) != 0 {
		switch setComm[0] {
		case "cd":
			if len(setComm) == 2 {
				if err := s.cd(setComm[1]); err != nil {
					fmt.Println(err)
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
					fmt.Fprintf(s.out, "%v ", err)
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
