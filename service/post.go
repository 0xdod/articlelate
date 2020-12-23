package service

import (
	"github.com/Kamva/mgm/v3"
	"github.com/fibreactive/articlelate/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostService interface {
	Get(filter interface{}) *models.Post
	GetByID(id interface{}) *models.Post
	GetByTitle(string) *models.Post
	GetAll() []*models.Post
	Filter(filter interface{}) []*models.Post
	Create(*models.Post) error
	Update(*models.Post) error
	Delete(id interface{}) error
}

//TODO
//Q: Do i really need to pass the data store around since i'm using mgm?
type PostStore struct{}

func NewPostStore() *PostStore {
	return &PostStore{}
}

func (ps *PostStore) Filter(filter interface{}) []*models.Post {
	var posts []*models.Post
	findOptions := options.Find().SetSort(bson.M{"created_at": -1})
	err := mgm.CollectionByName("posts").SimpleFind(&posts, filter, findOptions)
	if err != nil {
		return nil
	}
	return posts
}

func (ps *PostStore) GetAll() []*models.Post {
	var posts []*models.Post
	findOptions := options.Find().SetSort(bson.M{"created_at": -1})
	if err := mgm.CollectionByName("posts").SimpleFind(&posts, bson.M{}, findOptions); err != nil {
		return nil
	}
	return posts
}

func (*PostStore) query(filter interface{}) *models.Post {
	var post models.Post
	if err := mgm.Coll(&post).First(filter, &post); err != nil {
		return nil
	}
	return &post
}

func (a *PostStore) Get(filter interface{}) *models.Post {
	return a.query(filter)
}

func (ps *PostStore) Create(a *models.Post) error {
	return mgm.Coll(a).Create(a)
}

func (ps *PostStore) GetByID(id interface{}) *models.Post {
	var post models.Post
	if err := mgm.Coll(&post).FindByID(id, &post); err != nil {
		return nil
	}
	return &post
}
func (ps *PostStore) GetByTitle(t string) *models.Post {
	return ps.query(bson.M{"title": t})
}

func (ps *PostStore) Update(a *models.Post) error {
	return mgm.Coll(a).Update(a)
}

func (ps *PostStore) Delete(id interface{}) error {
	a := ps.GetByID(id)
	return mgm.Coll(a).Delete(a)
}
