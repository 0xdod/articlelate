package service

import (
	"bytes"
	"encoding/gob"
	"strings"

	"github.com/Kamva/mgm/v3"
	"github.com/fibreactive/articlelate/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

//pepper added to salted password
var userPwPepper = "secret-random-string-lolx"

type UserService interface {
	Authenticate(string, string) *models.User
	FindByToken(string) *models.User
	GenerateAuthToken(*models.User) (string, error)
	Get(string, interface{}) *models.User
	GetByID(id interface{}) *models.User
	Create(*models.User) error
	Update(*models.User) error
	Delete(id interface{}) error
	//TODO
	// AddFollower(*User, followerID interface{}) error
	// RemoveFollower(*User, followerID interface{}) error
	// IsFollower(userID, followerID interface{}) (bool, error)
}

type UserStore struct{}

func NewUserStore() *UserStore {
	return &UserStore{}
}

func (*UserStore) query(filter interface{}) *models.User {
	var user models.User
	if err := mgm.Coll(&user).First(filter, &user); err != nil {
		return nil
	}
	return &user
}

func (u *UserStore) GenerateAuthToken(user *models.User) (string, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(user)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (us *UserStore) FindByToken(token string) *models.User {
	dec := gob.NewDecoder(bytes.NewReader([]byte(token)))
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return nil
	}
	return &u
}

func (us *UserStore) Authenticate(login, password string) *models.User {
	user := us.Get("email", login)
	if user == nil {
		user = us.Get("username", login)
	}
	if user == nil {
		return nil
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash),
		[]byte(password+userPwPepper))
	if err != nil {
		// Invalid password
		return nil
	}
	return user
}

func (*UserStore) Create(u *models.User) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password+userPwPepper), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Email = strings.ToLower(u.Email)
	u.PasswordHash = string(hashedBytes)
	u.Password = ""
	return mgm.Coll(u).Create(u)
}

func (*UserStore) GetByID(id interface{}) *models.User {
	var user models.User
	if err := mgm.Coll(&user).FindByID(id, &user); err != nil {
		return nil
	}
	return &user
}

func (us *UserStore) Get(key string, value interface{}) *models.User {
	return us.query(bson.M{key: value})
}

func (*UserStore) Update(u *models.User) error {
	return mgm.Coll(u).Update(u)
}

func (us *UserStore) Delete(id interface{}) error {
	u := us.GetByID(id)
	return mgm.Coll(u).Delete(u)
}
