package service

import (
	"fmt"

	"github.com/puti-projects/puti/internal/pkg/constvar"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const (
	// B to byte
	B = 1
	// KB to B
	KB = 1024 * B
	// MB to KB
	MB = 1024 * KB
	// GB to MB
	GB = 1024 * MB
)

// SystemInfo system infomation
type SystemInfo struct {
	Hostname string `json:"hostname"`
	Uptime   string `json:"uptime"`
	OS       string `json:"os"`
	Platform string `json:"platform"`
}

// Health basic struct include health status
type Health struct {
	HealthStatus string `json:"healthStatus"`
}

// HealthPercent include health precentage
type HealthPercent struct {
	UsedPercent int `json:"usedPercent"`
}

// DiskHealth includes used disk, total disk, userd percent and health status
type DiskHealth struct {
	Health
	HealthPercent

	UsedMB  int `json:"usedMB"`
	UsedGB  int `json:"usedGB"`
	TotalMB int `json:"totalMB"`
	TotalGB int `json:"totalGB"`
}

// RAMHealth includes used RAM, total RAM, userd percent and health status
type RAMHealth struct {
	Health
	HealthPercent

	UsedMB  int `json:"usedMB"`
	UsedGB  int `json:"usedGB"`
	TotalMB int `json:"totalMB"`
	TotalGB int `json:"totalGB"`
}

// CPUHealth includes CPU cores, load average for one minutes, load average for five minutes, load average for fifteen minutes and health status
type CPUHealth struct {
	Health

	CPUCores      int     `json:"cpuCores"`
	LoadAverage1  float64 `json:"loadAverage1"`
	LoadAverage5  float64 `json:"loadAverage5"`
	LoadAverage15 float64 `json:"loadAverage15"`
}

// GetHealthStatusByPercent get health status by precentage
func (h *HealthPercent) GetHealthStatusByPercent() string {
	if h.UsedPercent >= 95 {
		return "CRITICAL"
	} else if h.UsedPercent >= 80 {
		return "WARNING"
	}

	return "NORMAL"
}

// GetHealthStatusByCores only fot CPU check; get health status by CPU cores
func (c *CPUHealth) GetHealthStatusByCores() string {
	if c.LoadAverage5 >= float64(c.CPUCores-1) {
		return "CRITICAL"
	} else if c.LoadAverage5 >= float64(c.CPUCores-2) {
		return "WARNING"
	}

	return "NORMAL"
}

// DiskCheck checks the disk usage.
func DiskCheck() *DiskHealth {
	var diskHealth *DiskHealth
	u, err := disk.Usage("/")
	if err != nil {
		diskHealth.HealthStatus = fmt.Sprintf("%s", err)
		return diskHealth
	}

	diskHealth = &DiskHealth{
		UsedMB:  int(u.Used) / MB,
		UsedGB:  int(u.Used) / GB,
		TotalMB: int(u.Total) / MB,
		TotalGB: int(u.Total) / GB,

		HealthPercent: HealthPercent{UsedPercent: int(u.UsedPercent)},
	}
	diskHealth.HealthStatus = diskHealth.GetHealthStatusByPercent()

	return diskHealth
}

// CPUCheck checks the cpu usage.
func CPUCheck() *CPUHealth {
	var cpuHealth *CPUHealth
	cores, err := cpu.Counts(false)
	if err != nil {
		cpuHealth.HealthStatus = fmt.Sprintf("%s", err)
		return cpuHealth
	}

	a, _ := load.Avg()

	cpuHealth = &CPUHealth{
		CPUCores:      cores,
		LoadAverage1:  a.Load1,
		LoadAverage5:  a.Load5,
		LoadAverage15: a.Load15,
	}
	cpuHealth.HealthStatus = cpuHealth.GetHealthStatusByCores()

	return cpuHealth
}

// RAMCheck checks the RAM usage.
func RAMCheck() *RAMHealth {
	var ramHealth *RAMHealth
	u, err := mem.VirtualMemory()
	if err != nil {
		ramHealth.HealthStatus = fmt.Sprintf("%s", err)
		return ramHealth
	}

	ramHealth = &RAMHealth{
		UsedMB:  int(u.Used) / MB,
		UsedGB:  int(u.Used) / GB,
		TotalMB: int(u.Total) / MB,
		TotalGB: int(u.Total) / GB,

		HealthPercent: HealthPercent{UsedPercent: int(u.UsedPercent)},
	}
	ramHealth.HealthStatus = ramHealth.GetHealthStatusByPercent()

	return ramHealth
}

// SystemInfoCheck get system information
func SystemInfoCheck() *SystemInfo {
	info, _ := host.Info()

	days := info.Uptime / constvar.SecondsPerDay
	daysSec := info.Uptime % constvar.SecondsPerDay
	hours := daysSec / constvar.SecondsPerHour
	hoursSec := daysSec % constvar.SecondsPerHour
	minutes := hoursSec / constvar.SecondsPerMinute

	var systemInfo *SystemInfo
	systemInfo = &SystemInfo{
		Hostname: info.Hostname,
		Uptime:   fmt.Sprintf("%d天 %d小时 %d分钟", days, hours, minutes),
		OS:       info.OS,
		Platform: info.Platform + " / " + info.PlatformVersion,
	}

	return systemInfo
}
