package paginator

import (
	"math"
)

type Pager struct {
	Pages       []int
	TotalPages  int
	FirstPage   int
	PrePage     int
	NextPage    int
	LastPage    int
	CurrentPage int
}

func Run(currentPage, everyPageNumber int, total int64, pagesNumber int) Pager {

	// 中间页位置
	middlePage := int(math.Ceil(float64(pagesNumber) / 2))
	if middlePage < 1 {
		middlePage = 1
	}

	// 首页
	var firstPage int = 1
	// 上一页
	var prePage int = 1
	// 下一页
	var nextPage int = 1
	// 尾页
	var lastPage int = 1

	// 根据总数和prepage每页数量 生成分页总数
	totalPages := int(math.Ceil(float64(total) / float64(everyPageNumber)))

	// 规划当前页
	if currentPage > totalPages {
		currentPage = totalPages
	}
	if currentPage <= 0 {
		currentPage = 1
	}

	// 中间页
	var pages []int
	switch {
	case currentPage > totalPages-pagesNumber && totalPages >= pagesNumber:
		start := totalPages - pagesNumber + 1
		pages = make([]int, pagesNumber)
		for i := range pages {
			pages[i] = start + i
		}
	case currentPage >= middlePage && totalPages >= pagesNumber:
		start := currentPage - 2
		pages = make([]int, pagesNumber)
		for i := range pages {
			pages[i] = start + i
		}
	default:
		pages = make([]int, int(math.Min(float64(pagesNumber), float64(totalPages))))
		for i := range pages {
			pages[i] = i + 1
		}
	}

	// 整理分页结果
	firstPage = 1
	if currentPage == 1 {
		firstPage = -1
	}
	prePage = currentPage - 1
	if prePage < 1 {
		prePage = -1
	}
	nextPage = currentPage + 1
	if nextPage >= totalPages {
		nextPage = -1
	}
	lastPage = totalPages
	if lastPage == currentPage {
		lastPage = -1
	}

	// 构建分页 map
	return Pager{
		TotalPages:  totalPages,
		FirstPage:   firstPage,
		PrePage:     prePage,
		Pages:       pages,
		NextPage:    nextPage,
		LastPage:    lastPage,
		CurrentPage: currentPage,
	}
}
