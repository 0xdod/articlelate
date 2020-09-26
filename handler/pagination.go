package handler

import (
	"errors"
	"math"

	"github.com/fibreactive/articlelate/models"
)

const minPage int = 1

var (
	PageNotAnInteger = errors.New("Pagination: Page not an integer")
	EmptyPage        = errors.New("Pagination: Page doesn't exist")
)

type Page struct {
	Items []*models.Article
	*Paginator
}

type Paginator struct {
	ObjectList   []*models.Article
	CurrentPage  int
	NextPage     int
	PreviousPage int
	Limit        int
	MaxPage      int
}

func NewPaginator(objectList []*models.Article, limit int) *Paginator {
	if objectList == nil {
		return nil
	}
	var maxPage = int(math.Ceil(float64(len(objectList)) / float64(limit)))
	return &Paginator{
		MaxPage:     maxPage,
		ObjectList:  objectList,
		Limit:       limit,
		CurrentPage: minPage,
	}
}

func (pg *Paginator) Page(page int) (*Page, error) {
	if page < minPage || page > pg.MaxPage {
		return nil, EmptyPage
	}
	lowerLimit := (page - 1) * pg.Limit
	upperLimit := lowerLimit + pg.Limit
	if len(pg.ObjectList) < upperLimit {
		upperLimit = len(pg.ObjectList)
	}
	pg.CurrentPage = page
	pg.NextPage = page + 1
	pg.PreviousPage = page - 1
	items := pg.ObjectList[lowerLimit:upperLimit]
	return &Page{items, pg}, nil
}

func (p *Page) CurrentPage() int {
	return p.Paginator.CurrentPage
}

func (p *Page) PreviousPage() int {
	return p.Paginator.PreviousPage
}

func (p *Page) NextPage() int {
	return p.Paginator.NextPage
}

func (p *Page) HasPrevious() bool {
	prev := p.PreviousPage()
	curr := p.CurrentPage()
	return prev >= minPage && prev < curr
}

func (p *Page) HasNext() bool {
	next := p.NextPage()
	curr := p.CurrentPage()
	return next <= p.Paginator.MaxPage && next > curr
}
