package service

import (
	"fmt"
	"github.com/puti-projects/puti/internal/pkg/constvar"
	"github.com/puti-projects/puti/internal/pkg/logger"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"strconv"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
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

const (
	statusERROR = "ERROR"
	statusCRITICAL = "CRITICAL"
	statusWARNING = "WARNING"
	statusNORMAL = "NORMAL"
)

// SystemInfo system information
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

// HealthPercent include health percentage
type HealthPercent struct {
	UsedPercent int `json:"usedPercent"`
}

// DiskHealth includes used disk, total disk, used percent and health status
type DiskHealth struct {
	Health
	HealthPercent

	Used  string `json:"used"`
	Total string `json:"total"`
}

// RAMHealth includes used RAM, total RAM, used percent and health status
type RAMHealth struct {
	Health
	HealthPercent

	Used  string `json:"used"`
	Total string `json:"total"`
}

// CPUHealth includes CPU cores, load average for one minutes, load average for five minutes, load average for fifteen minutes and health status
type CPUHealth struct {
	Health

	CPUCores      int     `json:"cpuCores"`
	LoadAverage1  float64 `json:"loadAverage1"`
	LoadAverage5  float64 `json:"loadAverage5"`
	LoadAverage15 float64 `json:"loadAverage15"`
}

// GetHealthStatusByPercent get health status by percentage
func (h *HealthPercent) GetHealthStatusByPercent() string {
	if h.UsedPercent == 0 {
		return statusERROR
	}

	if h.UsedPercent >= 95 {
		return statusCRITICAL
	} else if h.UsedPercent >= 80 {
		return statusWARNING
	}

	return statusNORMAL
}

// GetHealthStatusByCores only fot CPU check; get health status by CPU cores
func (c *CPUHealth) GetHealthStatusByCores() string {
	var criticalLevel, warningLevel float64
	switch c.CPUCores {
	case 1:
		criticalLevel = float64(c.CPUCores) - 0.1
		warningLevel = float64(c.CPUCores) - 0.25
		break
	case 2:
		criticalLevel = float64(c.CPUCores) - 0.25
		warningLevel = float64(c.CPUCores) - 0.5
		break
	default:
		criticalLevel = float64(c.CPUCores-1)
		warningLevel = float64(c.CPUCores-2)
		break
	}

	if c.LoadAverage5 >= criticalLevel {
		return statusCRITICAL
	} else if c.LoadAverage5 >= warningLevel {
		return statusWARNING
	}

	return statusNORMAL
}

// DiskCheck checks the disk usage.
func DiskCheck() *DiskHealth {
	var diskHealth *DiskHealth
	u, err := disk.Usage("/")
	if err != nil {
		logger.Errorf("error when getting disk info: %s", err)

		diskHealth = &DiskHealth{
			Health: Health{HealthStatus: statusERROR},
		}
		return diskHealth
	}

	diskHealth = &DiskHealth{
		Used:  transSize(u.Used),
		Total: transSize(u.Total),
		HealthPercent: HealthPercent{UsedPercent: int(u.UsedPercent)},
	}
	diskHealth.HealthStatus = diskHealth.GetHealthStatusByPercent()

	return diskHealth
}

// RAMCheck checks the RAM usage.
func RAMCheck() *RAMHealth {
	var ramHealth *RAMHealth

	u, err := mem.VirtualMemory()
	if err != nil {
		logger.Errorf("error when getting RAM info: %s", err)

		ramHealth = &RAMHealth{
			HealthPercent: HealthPercent{UsedPercent: 0},
		}
		ramHealth.HealthStatus = ramHealth.GetHealthStatusByPercent()
		return ramHealth
	}

	ramHealth = &RAMHealth{
		Used:  transSize(u.Used),
		Total: transSize(u.Total),
		HealthPercent: HealthPercent{UsedPercent: int(u.UsedPercent)},
	}
	ramHealth.HealthStatus = ramHealth.GetHealthStatusByPercent()

	return ramHealth
}

// CPUCheck checks the cpu usage.
func CPUCheck() *CPUHealth {
	var cpuHealth *CPUHealth
	cores, err := cpu.Counts(false)
	if err != nil {
		logger.Errorf("error when getting cpu info: %s", err)

		cpuHealth = &CPUHealth{
			Health: Health{HealthStatus: statusERROR},
		}
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

// transSize tool function for change size
func transSize(size uint64) string {
	var transfer string
	if size >= 1 * GB {
		sizeFloat, _ := strconv.ParseFloat(fmt.Sprintf("%d.%.2d", size / GB, size % GB), 64)
		sizeGB := strconv.FormatFloat(sizeFloat,'f',2,64)
		transfer = fmt.Sprintf("%s GB", sizeGB)
	} else {
		transfer = fmt.Sprintf("%d MB", size / MB)
	}

	return transfer
}
