package statistics

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"

	"github.com/gin-gonic/gin"
)

// SystemInfo system info for response
type SystemInfo struct {
	Info *service.SystemInfo `json:"info"`
	Disk *service.DiskHealth `json:"disk"`
	RAM  *service.RAMHealth  `json:"ram"`
	CPU  *service.CPUHealth  `json:"cpu"`
}

// System get system info
func System(c *gin.Context) {
	systemInfo := &SystemInfo{
		Info: service.SystemInfoCheck(),
		Disk: service.DiskCheck(),
		RAM:  service.RAMCheck(),
		CPU:  service.CPUCheck(),
	}

	Response.SendResponse(c, nil, systemInfo)
}
