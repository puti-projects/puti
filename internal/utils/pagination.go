package utils

import "math"

// Page pagination
type Page struct {
	ShowSize int `json:"showSize"`
	PageSize int `json:"pageSize"`

	CurrentPage int `json:"currentPage"`
	TotalPage   int `json:"totalPage"`
	TotalRecord int `json:"totalRecord"`

	NextPageNum     int   `json:"nextPageNum"`
	PreviousPageNum int   `json:"previousPageNum"`
	FirstPageNum    int   `josn:"firstPageNum"`
	LastPageNum     int   `json:"lastPageNum"`
	PageNums        []int `json:"pageNums"`
}

// Pagination include page and page url
type Pagination struct {
	Page    *Page
	PageURL string `json:"pageURL"`
}

const (
	defaultPageSize = 10
	defaultShowSize = 5
	defaultPageURL  = ""
)

// SetPageURL set page url, default is empty string
func (p *Pagination) SetPageURL(url string) {
	p.PageURL = url
}

// GetPagination return pagiantion and default page pre url
func GetPagination(totalRecord, currentPage, pageSizeIn, showSizeIn int) *Pagination {
	page := newPagination(totalRecord, currentPage, pageSizeIn, showSizeIn)

	return &Pagination{
		Page:    page,
		PageURL: defaultPageURL,
	}
}

// newPagination return pagiantion
func newPagination(totalRecord, currentPage, pageSizeIn, showSizeIn int) *Page {
	pageSize, showSize := paginationSetting(pageSizeIn, showSizeIn)

	totalPage := int(math.Ceil(float64(totalRecord) / float64(pageSize)))

	previousPageNum := currentPage - 1
	if 1 > previousPageNum {
		previousPageNum = 0
	}
	nextPageNum := currentPage + 1
	if nextPageNum > totalPage {
		nextPageNum = 0
	}

	var pageNums []int
	if totalPage < showSize {
		for i := 0; i < totalPage; i++ {
			pageNums = append(pageNums, i+1)
		}
	} else {
		first := currentPage + 1 - showSize/2
		if first < 1 {
			first = 1
		}
		if first+showSize > totalPage {
			first = totalPage - showSize + 1
		}
		for i := 0; i < showSize; i++ {
			pageNums = append(pageNums, first+i)
		}
	}

	firstPageNum := 0
	lastPageNum := 0
	if 0 < len(pageNums) {
		firstPageNum = pageNums[0]
		lastPageNum = pageNums[len(pageNums)-1]
	}

	return &Page{
		ShowSize:        showSize,
		PageSize:        pageSize,
		CurrentPage:     currentPage,
		TotalPage:       totalPage,
		TotalRecord:     totalRecord,
		PageNums:        pageNums,
		PreviousPageNum: previousPageNum,
		NextPageNum:     nextPageNum,
		FirstPageNum:    firstPageNum,
		LastPageNum:     lastPageNum,
	}
}

func paginationSetting(pageSize, showSize int) (int, int) {
	if pageSize == 0 {
		pageSize = defaultPageSize
	}

	if showSize == 0 {
		showSize = defaultShowSize
	}

	return pageSize, showSize
}
