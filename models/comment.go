package models

import (
	"github.com/Kamva/mgm/v3"
)

type Comment struct {
	mgm.DefaultModel `bson:",inline"`
	Post             *Post    `bson:"post"`
	Author           *User    `bson:"author"`
	Content          string   `bson:"content"`
	Likes            []string `bson:"likes"`
}

func NewComment(a *User, p *Post, c string) *Comment {
	return &Comment{
		Post:    p,
		Author:  a,
		Content: c,
	}
}
