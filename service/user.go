package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Kamva/mgm/v3"
	"github.com/dgrijalva/jwt-go"
	"github.com/fibreactive/articlelate/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//pepper added to salted password
var userPwPepper = "secret-random-string-lolx"

type UserService interface {
	Authenticate(string, string) *models.User
	GenerateAuthToken(*models.User) (string, error)
	FindByToken(string) *models.User
	RemoveAuthToken(string) error
	GetByID(id interface{}) *models.User
	GetByEmail(string) *models.User
	GetByUsername(string) *models.User
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

func (us *UserStore) FindByToken(tokenString string) *models.User {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("Hello"), nil
	})
	if err != nil {
		return nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil
	}
	id, err := primitive.ObjectIDFromHex(claims["_id"].(string))
	if err != nil {
		return nil
	}
	return us.query(bson.M{"_id": id, "tokens.access": claims["access"].(string), "tokens.token": tokenString})
}

func (u *UserStore) GenerateAuthToken(user *models.User) (string, error) {
	access := "auth"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":    user.ID.Hex(),
		"access": access,
	})
	//move to env
	tokenString, err := token.SignedString([]byte("Hello"))
	if err != nil {
		return "", err
	}
	userToken := &models.Token{
		Access: access,
		Token:  tokenString,
	}
	if user.Tokens != nil {
		for _, v := range user.Tokens {
			if v.Access == userToken.Access && v.Token == userToken.Token {
				user.Tokens = []*models.Token{userToken}
			}
		}
	} else {
		user.Tokens = append(user.Tokens, userToken)
	}

	if err := u.Update(user); err != nil {
		return "", err
	}
	return tokenString, nil
}

func (u *UserStore) RemoveAuthToken(t string) error {
	user := u.FindByToken(t)
	if user == nil {
		return errors.New("Error finding user")
	}
	user.Tokens = nil
	return u.Update(user)
}

func (us *UserStore) Authenticate(login, password string) *models.User {
	user := us.GetByEmail(login)
	if user == nil {
		user = us.GetByUsername(login)
	}
	if user == nil {
		return nil
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password+userPwPepper))
	if err != nil {
		// Invalid password
		return nil
	}
	return user
}

func (*UserStore) query(filter interface{}) *models.User {
	var user models.User
	if err := mgm.Coll(&user).First(filter, &user); err != nil {
		return nil
	}
	return &user
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
func (us *UserStore) GetByEmail(e string) *models.User {
	return us.query(bson.M{"email": e})

}
func (us *UserStore) GetByUsername(u string) *models.User {
	return us.query(bson.M{"username": u})
}

func (*UserStore) Update(u *models.User) error {
	return mgm.Coll(u).Update(u)
}

func (us *UserStore) Delete(id interface{}) error {
	u := us.GetByID(id)
	return mgm.Coll(u).Delete(u)
}
