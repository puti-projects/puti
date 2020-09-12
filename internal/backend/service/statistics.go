package service

import (
	"github.com/puti-projects/puti/internal/backend/dao"
)

// DashboardData some statistics index
type DashboardData struct {
	TotalViews    int64
	TotalComments int64
	TotalArticles int64
	TotalMedia    int64
}

// GetDashboardStatisticsData get dashboard statistics data
// TODO store in cache first
func GetDashboardStatisticsData() (*DashboardData, error) {
	dashboardData := &DashboardData{
		TotalViews:    dao.Engine.GetPostTotalView(),
		TotalComments: 0, // TODO comment Features
		TotalArticles: dao.Engine.GetTotalArticles(),
		TotalMedia:    dao.Engine.GetTotalMedia(),
	}

	return dashboardData, nil
}
