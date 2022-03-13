package shell

import "errors"

var (
	errChangeDir = errors.New("the command must have 1 argument")
	errPWD       = errors.New("the command must not have arguments")
	errKill      = errors.New("the command must have at least 1 argument")
	errEcho      = errors.New("the command must have at least 1 argument")
	errPS        = errors.New("the command must not have arguments")
	errExec      = errors.New("the command must have at least 1 argument")
)
