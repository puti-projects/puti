package statistics

import (
	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/pkg/system"

	"github.com/gin-gonic/gin"
)

// SystemInfo system info for response
type SystemInfo struct {
	Info *system.SystemInfo `json:"info"`
	Disk *system.DiskHealth `json:"disk"`
	RAM  *system.RAMHealth  `json:"ram"`
	CPU  *system.CPUHealth  `json:"cpu"`
}

// System get system info
func System(c *gin.Context) {
	systemInfo := &SystemInfo{
		Info: system.SystemInfoCheck(),
		Disk: system.DiskCheck(),
		RAM:  system.RAMCheck(),
		CPU:  system.CPUCheck(),
	}

	api.SendResponse(c, nil, systemInfo)
}
