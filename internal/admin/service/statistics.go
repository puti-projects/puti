package service

// DashboardData some statistics index
type DashboardData struct {
	TotalViews    int64
	TotalComments int64
	TotalArticles int64
	TotalMedia    int64
}

// GetDashboardStatisticsData get dashboard statistics data
// TODO store in cache first
func (svc Service) GetDashboardStatisticsData() (*DashboardData, error) {
	dashboardData := &DashboardData{
		TotalViews:    svc.dao.GetPostTotalView(),
		TotalComments: 0, // TODO comment Features
		TotalArticles: svc.dao.GetTotalArticles(),
		TotalMedia:    svc.dao.GetTotalMedia(),
	}

	return dashboardData, nil
}
