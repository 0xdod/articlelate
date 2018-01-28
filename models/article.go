package models

import (
	"fmt"

	"github.com/Kamva/mgm/v3"
	"github.com/gosimple/slug"
)

type Article struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string     `json:"title" bson:"title"`
	Content          string     `json:"content" bson:"content"`
	Likes            []string   `json:"likes" bson:"likes"`
	Comments         []*Comment `json:"comments" bson:"-"`
	Author           *User      `json:"author" bson:"author"`
	Slug             string     `json:"slug" bson:"slug"`
}

func NewArticle(author *User, title, body string) *Article {
	return &Article{
		Title:   title,
		Content: body,
		Author:  author,
	}
}

func (a *Article) GetAbsoluteURL() string {
	return fmt.Sprintf("/p/%s/%s", a.Author.Username, a.Slug)
}

func (a *Article) Created() error {
	id := a.ID.Hex()
	suffix := fmt.Sprintf(" %s", id[len(id)-4:])
	a.Slug = slug.Make(a.Title + suffix)
	return mgm.Coll(a).Update(a)
}
