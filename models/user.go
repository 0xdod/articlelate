package models

import (
	"github.com/Kamva/mgm/v3"
)

// revise user authentication flow
type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string   `bson:"username"`
	Email            string   `bson:"email"`
	Followers        int      `bson:"followers"`
	Following        int      `bson:"following"`
	Password         string   `bson:"-"`
	PasswordHash     string   `bson:"passwordHash"`
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

func (u *User) OwnsPost(p *Post) bool {
	return u.Username == p.Author.Username
}

func (u *User) LikedPost(p *Post) bool {
	for _, username := range p.Likes {
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

	case *Post:
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
	case *Post:
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
