package models

import (
	"github.com/Kamva/mgm/v3"
)

type User struct {
	// DefaultModel add _id,created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Username         string   `json:"username" bson:"username"`
	Email            string   `json:"email" bson:"email"`
	Followers        int      `json:"followers" bson:"followers"`
	Following        int      `json:"following" bson:"following"`
	Password         string   `json:"password" bson:"-"`
	PasswordHash     string   `json:"passwordHash" bson:"passwordHash"`
	Tokens           []*Token `bson:"tokens"`
}

type Token struct {
	Access string `bson:"access" binding:required`
	Token  string `bson:"token" binding:required`
}

func NewUser(name, email, pswd string) *User {
	return &User{
		Username: name,
		Email:    email,
		Password: pswd,
	}
}

func (u *User) OwnsArticle(a *Article) bool {
	return u.Username == a.Author.Username
}

func (u *User) LikedArticle(a *Article) bool {
	for _, username := range a.Likes {
		if username == u.Username {
			return true
		}
	}
	return false
}

func (u *User) Like(obj interface{}) {
	switch obj := obj.(type) {
	case *Comment:
		for _, v := range obj.Likes {
			if v == u.Username {
				return
			}
		}
		obj.Likes = append(obj.Likes, u.Username)

	case *Article:
		for _, v := range obj.Likes {
			if v == u.Username {
				return
			}
		}
		obj.Likes = append(obj.Likes, u.Username)
	}
}

func (u *User) Unlike(obj interface{}) {
	switch obj := obj.(type) {
	case *Comment:
		for i, v := range obj.Likes {
			if v == u.Username {
				firstHalf := obj.Likes[:i]
				secondHalf := obj.Likes[i+1:]
				fullSlice := append(firstHalf, secondHalf...)
				obj.Likes = fullSlice
				break
			}
		}
	case *Article:
		for i, v := range obj.Likes {
			if v == u.Username {
				firstHalf := obj.Likes[:i]
				secondHalf := obj.Likes[i+1:]
				fullSlice := append(firstHalf, secondHalf...)
				obj.Likes = fullSlice
				break
			}
		}
	}

}
