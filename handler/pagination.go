package handler

import (
	"errors"
	"math"
)

const (
	DefaultLimit  = 10
	MinPageNumber = 1
)

var (
	EmptyPage = errors.New("Pagination: Page doesn't exist")
)

type Adapter interface {
	Slice(offset, limit int, data interface{}) error
	Counts() int64
}

type Page struct {
	*Paginator
	Items interface{}
}

type Paginator struct {
	adapter Adapter
	limit   int
	page    int
}

func NewPaginator(a Adapter, limit int) *Paginator {
	if limit < 1 {
		limit = DefaultLimit
	}
	return &Paginator{
		adapter: a,
		limit:   limit,
	}
}

func (p *Paginator) Result(data interface{}) error {
	if p.page < MinPageNumber || p.page > p.PageCount() {
		return EmptyPage
	}
	skip := (p.page - 1) * p.limit
	return p.adapter.Slice(skip, p.limit, data)
}

func (p *Paginator) Page() int {
	return p.page
}

func (p *Paginator) SetPage(page int) {
	p.page = page
}

func (p *Paginator) PreviousPage() int {
	prevPage := p.page - 1
	if prevPage < MinPageNumber {
		return MinPageNumber
	}
	return prevPage
}

func (p *Paginator) NextPage() int {
	nextPage := p.page + 1
	maxPage := p.PageCount()
	if nextPage > maxPage {
		return maxPage
	}
	return nextPage
}

func (p *Paginator) HasPrevious() bool {
	prev := p.PreviousPage()
	curr := p.Page()
	return prev >= MinPageNumber && prev < curr
}

func (p *Paginator) HasNext() bool {
	next := p.NextPage()
	curr := p.Page()
	return next <= p.PageCount() && next > curr
}

func (p *Paginator) PageCount() int {
	n := int(math.Ceil(float64(p.adapter.Counts()) / float64(p.limit)))
	if n < 1 {
		return 1
	}
	return n
}
