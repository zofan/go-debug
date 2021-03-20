package debug

import (
	"net"
	"os"
	"runtime"
)

type System struct {
	Arch   string
	OS     string
	NumCPU int

	Hostname  string
	IpList    []net.IP
	PageSize  int
	TempDir   string
	CacheDir  string
	ConfigDir string
	HomeDir   string
}

func SystemInfo() *System {
	info := &System{}

	info.Arch = runtime.GOARCH
	info.OS = runtime.GOOS
	info.NumCPU = runtime.NumCPU()
	info.Hostname, _ = os.Hostname()
	info.IpList, _ = net.LookupIP(info.Hostname)
	info.PageSize = os.Getpagesize()

	info.TempDir = os.TempDir()
	info.CacheDir, _ = os.UserCacheDir()
	info.ConfigDir, _ = os.UserConfigDir()
	info.HomeDir, _ = os.UserHomeDir()

	return info
}
