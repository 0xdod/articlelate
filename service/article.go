package service

import (
	"github.com/Kamva/mgm/v3"
	"github.com/fibreactive/articlelate/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ArticleService interface {
	GetAll() []*models.Article
	GetByID(id interface{}) *models.Article
	GetByTitle(string) *models.Article
	Create(*models.Article) error
	Update(*models.Article) error
	Delete(id interface{}) error
}

//TODO
//Q: Do i really need to pass the data store around since i'm using mgm?
type ArticleStore struct{}

func NewArticleStore() *ArticleStore {
	return &ArticleStore{}
}
func (as *ArticleStore) GetAll() []*models.Article {
	var articles []*models.Article
	findOptions := options.Find().SetSort(bson.M{"created_at": -1})
	if err := mgm.CollectionByName("articles").SimpleFind(&articles, bson.M{}, findOptions); err != nil {
		return nil
	}
	return articles
}

func (*ArticleStore) query(filter interface{}) *models.Article {
	var article models.Article
	if err := mgm.Coll(&article).First(filter, &article); err != nil {
		return nil
	}
	return &article
}

func (as *ArticleStore) Create(a *models.Article) error {
	return mgm.Coll(a).Create(a)
}

func (as *ArticleStore) GetByID(id interface{}) *models.Article {
	var article models.Article
	if err := mgm.Coll(&article).FindByID(id, &article); err != nil {
		return nil
	}
	return &article
}
func (as *ArticleStore) GetByTitle(t string) *models.Article {
	return as.query(bson.M{"title": t})
}

func (as *ArticleStore) Update(a *models.Article) error {
	return mgm.Coll(a).Update(a)
}

func (as *ArticleStore) Delete(id interface{}) error {
	a := as.GetByID(id)
	return mgm.Coll(a).Delete(a)
}
