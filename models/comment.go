package models

import (
	"github.com/Kamva/mgm/v3"
)

type Comment struct {
	mgm.DefaultModel `bson:",inline"`
	Article          *Article `json:"article" bson:"article"`
	Author           *User    `json:"author" bson:"author"`
	Content          string   `json:"content" bson:"content"`
	Likes            []string `json:"content" bson:"likes"`
}

func NewComment(author *User, article *Article, content string) *Comment {
	return &Comment{
		Article: article,
		Author:  author,
		Content: content,
	}
}
