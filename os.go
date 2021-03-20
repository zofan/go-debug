package debug

import (
	"os"
)

type OS struct {
	WorkDirectory string
	Args          []string
	EGID          int
	EUID          int
	GID           int
	PID           int
	PPID          int
	UID           int
}

func OsInfo() *OS {
	info := &OS{}

	info.WorkDirectory, _ = os.Getwd()
	info.Args = os.Args
	//info[`env`] = os.Environ()
	info.EGID = os.Getegid()
	info.EUID = os.Geteuid()
	info.GID = os.Getgid()
	info.PID = os.Getpid()
	info.PPID = os.Getppid()
	info.UID = os.Getuid()

	return info
}
