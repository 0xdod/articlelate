package models

import (
	"fmt"

	mgm "github.com/Kamva/mgm/v3"
	"github.com/gosimple/slug"
)

type Post struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string     `bson:"title"`
	Content          string     `bson:"content"`
	Likes            []string   `bson:"likes"`
	Comments         []*Comment `bson:"-"`
	Author           *User      `bson:"author"`
	Slug             string     `bson:"slug"`
	Modified         bool       `bson:"modified"`
}

func NewPost(author *User, title, body string) *Post {
	return &Post{
		Title:   title,
		Content: body,
		Author:  author,
	}
}

func (p *Post) GetAbsoluteURL() string {
	return fmt.Sprintf("/p/%s/%s", p.Author.Username, p.Slug)
}

func (p *Post) Created() error {
	id := p.ID.Hex()
	suffix := fmt.Sprintf(" %s", id[len(id)-4:])
	p.Slug = slug.Make(p.Title + suffix)
	return mgm.Coll(p).Update(p)
}
