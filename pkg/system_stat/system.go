package system_stat

import (
	"runtime"
)

type SystemInfo struct {
	NumCpu int32
	Memory int32
}

func StatCpuMemory() *SystemInfo {
	numCPU := runtime.NumCPU()
	// 获取内存信息
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	totalMemory := memStats.TotalAlloc / (1024 * 1024) // 将字节转换为 MB
	return &SystemInfo{
		NumCpu: int32(numCPU),
		Memory: int32(totalMemory),
	}
}
