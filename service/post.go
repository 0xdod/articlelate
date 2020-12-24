package service

import (
	"context"
	"log"
	"time"

	"github.com/Kamva/mgm/v3"
	"github.com/fibreactive/articlelate/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostService interface {
	Get(filter interface{}) *models.Post
	GetByID(id interface{}) *models.Post
	All(interface{}) error
	Filter(interface{}, interface{}) error
	Create(*models.Post) error
	Update(*models.Post) error
	Delete(id interface{}) error
}

func PopulateIndex() {
	CreateSlugIndex()
	CreateTextIndex()
	log.Println("Successfully create the indexes")
}

func CreateTextIndex() {
	coll := mgm.Coll(&models.Post{}).Collection
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	index := mongo.IndexModel{}
	index.Keys = bson.D{{"title", "text"}, {"content", "text"}}
	index.Options = options.Index().SetBackground(true)
	coll.Indexes().CreateOne(context.Background(), index, opts)
}

func CreateSlugIndex() {
	coll := mgm.Coll(&models.Post{}).Collection
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	index := mongo.IndexModel{}
	index.Keys = bson.D{{"slug", 1}, {"author.username", 1}}
	index.Options = options.Index().SetBackground(true).SetUnique(true)
	coll.Indexes().CreateOne(context.Background(), index, opts)
}

//TODO
//Q: Do i really need to pass the data store around since i'm using mgm?
type PostMongo struct{}

type MongoAdapter struct {
	filter interface{}
}

func NewPostMongo() *PostMongo {
	PopulateIndex()
	return &PostMongo{}
}

func NewMongoAdapter() *MongoAdapter {
	return &MongoAdapter{bson.M{}}
}

func (ma *MongoAdapter) Slice(offset, limit int, data interface{}) error {
	findOptions := options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)).
		SetSort(bson.M{"created_at": -1})
	return mgm.Coll(&models.Post{}).SimpleFind(data, ma.filter, findOptions)
}

func (ma *MongoAdapter) SetFilter(filter interface{}) {
	ma.filter = filter
}

func (ma *MongoAdapter) Counts() int64 {
	nums, err := mgm.Coll(&models.Post{}).CountDocuments(context.Background(), ma.filter)
	if err != nil {
		return 0
	}
	return nums
}

func (pm *PostMongo) Filter(posts, filter interface{}) error {
	findOptions := options.Find().SetSort(bson.M{"created_at": -1})
	return mgm.Coll(&models.Post{}).SimpleFind(posts, filter, findOptions)
}

func (pm *PostMongo) All(posts interface{}) error {
	return pm.Filter(posts, bson.M{})
}

func (*PostMongo) query(filter interface{}) *models.Post {
	var post models.Post
	if err := mgm.Coll(&post).First(filter, &post); err != nil {
		return nil
	}
	return &post
}

func (pm *PostMongo) Get(filter interface{}) *models.Post {
	return pm.query(filter)
}

func (pm *PostMongo) GetByID(id interface{}) *models.Post {
	post := &models.Post{}
	err := mgm.Coll(post).FindByID(id, post)
	if err != nil {
		return nil
	}
	return post
}

func (pm *PostMongo) Create(a *models.Post) error {
	return mgm.Coll(a).Create(a)
}

func (pm *PostMongo) Update(a *models.Post) error {
	return mgm.Coll(a).Update(a)
}

func (pm *PostMongo) Delete(id interface{}) error {
	a := pm.GetByID(id)
	return mgm.Coll(a).Delete(a)
}
