package models

import (
	"github.com/Kamva/mgm/v3"
)

type Article struct {
	// DefaultModel add _id,created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Title            string     `json:"title" bson:"title"`
	Content          string     `json:"content" bson:"content"`
	Likes            []string   `json:"likes" bson:"likes"`
	Comments         []*Comment `json:"comments" bson:"-"`
	Author           *User      `json:"author" bson:"author"`
}

func NewArticle(author *User, title, body string) *Article {
	return &Article{
		Title:   title,
		Content: body,
		Author:  author,
	}
}
