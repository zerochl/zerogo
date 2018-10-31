package db

import (
	"zerogo/utils"
	"math"
	"net/url"
	"strconv"
)

const DEFAULT_PER_PAGE = 10
const MAX_SHOW_PAGE = 9

type Mode int

const (
	FULL Mode = 1 + iota
	NEXT_ONLY
)

type Pagination struct {
	Page      int
	PerPage   int
	Total     int
	Data      []interface{}
	hasNext   bool
	pageRange []int
	url       *url.URL
}

func NewPagination(page int, total int, hasNext bool) *Pagination {
	pagination := new(Pagination)

	if page <= 0 {
		page = 1
	}
	pagination.PerPage = DEFAULT_PER_PAGE
	pagination.Page = page
	pagination.Total = total
	pagination.hasNext = hasNext
	return pagination
}

func (p *Pagination) setPerPage(perPage int) {
	p.PerPage = perPage
}

func (p *Pagination) TotalPages() int {
	return (p.Total + p.PerPage - 1) / p.PerPage;
}

func (p *Pagination) NextPage() int {
	if (p.Page < p.TotalPages()) {
		return p.Page + 1;
	}
	return -1;
}

func (p *Pagination) PrevPage() int {
	if p.Page <= 1 {
		return -1
	} else {
		return p.Page - 1
	}
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.PerPage + 1;
}

func ( p *Pagination) HasNext() bool {
	return p.hasNext
}

func (p *Pagination) SetData(container interface{}) {
	p.Data = utils.ToSlice(container)
}

func (p *Pagination) Pages(maxShowPages int) []int {

	if (maxShowPages < 5 || maxShowPages > MAX_SHOW_PAGE) {
		maxShowPages = MAX_SHOW_PAGE;
	}
	middlePageNum := maxShowPages / 2
	if p.pageRange == nil && p.Total > 0 {
		var pages []int
		pageNums := p.TotalPages()
		page := p.Page
		switch {
		case page >= pageNums - middlePageNum && pageNums > maxShowPages:
			start := pageNums - maxShowPages + 1
			pages = make([]int, maxShowPages)
			for i := range pages {
				pages[i] = start + i
			}
		case page >= (middlePageNum + 1) && pageNums > maxShowPages:
			start := page - middlePageNum
			pages = make([]int, int(math.Min(float64(maxShowPages), float64(page + middlePageNum + 1))))
			for i := range pages {
				pages[i] = start + i
			}
		default:
			pages = make([]int, int(math.Min(float64(maxShowPages), float64(pageNums))))
			for i := range pages {
				pages[i] = i + 1
			}
		}
		p.pageRange = pages
	}
	return p.pageRange
}

func (p *Pagination) PageLink(page int) string {
	values := p.url.Query()
	values.Set("page", strconv.Itoa(page))
	p.url.RawQuery = values.Encode()
	return p.url.String()
}

// Returns URL to the previous page.
func (p *Pagination) PageLinkPrev() (link string) {
	if p.HasPrev() {
		link = p.PageLink(p.Page - 1)
	}
	return
}

// Returns URL to the next page.
func (p *Pagination) PageLinkNext() (link string) {
	if p.HasNext() {
		link = p.PageLink(p.Page + 1)
	}
	return
}

// Returns URL to the first page.
func (p *Pagination) PageLinkFirst() (link string) {
	return p.PageLink(1)
}

// Returns URL to the last page.
func (p *Pagination) PageLinkLast() (link string) {
	return p.PageLink(p.TotalPages())
}

func ( p *Pagination) HasPrev() bool {
	return p.Page > 1
}

func (p *Pagination) IsActive(pagea int) bool {
	return p.Page == pagea
}

func (p *Pagination) SetUrl(url *url.URL) {
	p.url = url
}









