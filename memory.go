package debug

import (
	"runtime"
)

type Runtime struct {
	Version      string
	Compiler     string
	GoRoot       string
	GoMaxProc    int
	NumGoroutine int
	NumCgoCall   int64
}

func RuntimeInfo() *Runtime {
	info := &Runtime{}

	info.Version = runtime.Version()
	info.Compiler = runtime.Compiler
	info.GoRoot = runtime.GOROOT()
	info.GoMaxProc = runtime.GOMAXPROCS(0)
	info.NumGoroutine = runtime.NumGoroutine()
	info.NumCgoCall = runtime.NumCgoCall()

	return info
}
