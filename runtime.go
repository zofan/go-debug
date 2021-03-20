package debug

import (
	"runtime"
)

type Memory struct {
	MallocMb     float32
	AllocMb      float32
	TotalAllocMb float32
	HeapAllocMb  float32
	HeapInuseMb  float32
	FreesMb      float32
	SysMb        float32
	NumGC        uint32
}

func MemoryInfo() *Memory {
	info := &Memory{}

	ms := &runtime.MemStats{}
	runtime.ReadMemStats(ms)

	info.MallocMb = float32(ms.Mallocs) / 1024 / 1024
	info.AllocMb = float32(ms.Alloc) / 1024 / 1024
	info.TotalAllocMb = float32(ms.TotalAlloc) / 1024 / 1024
	info.HeapAllocMb = float32(ms.HeapAlloc) / 1024 / 1024
	info.HeapInuseMb = float32(ms.HeapInuse) / 1024 / 1024
	info.FreesMb = float32(ms.Frees) / 1024 / 1024
	info.SysMb = float32(ms.Sys) / 1024 / 1024
	info.NumGC = ms.NumGC

	return info
}
