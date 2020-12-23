package service

import (
	"github.com/Kamva/mgm/v3"
	"github.com/fibreactive/articlelate/models"
	"go.mongodb.org/mongo-driver/bson"
)

type CommentService interface {
	GetByPost(*models.Post) []*models.Comment
	GetByID(id interface{}) *models.Comment
	Create(*models.Comment) error
	Update(*models.Comment) error
	Delete(id interface{}) error
}

type CommentStore struct{}

func NewCommentStore() *CommentStore {
	return &CommentStore{}
}

func (*CommentStore) query(filter interface{}) *models.Comment {
	var comment models.Comment
	if err := mgm.Coll(&comment).First(filter, &comment); err != nil {
		return nil
	}
	return &comment
}

func (cs *CommentStore) GetByPost(p *models.Post) []*models.Comment {
	var comments []*models.Comment
	if err := mgm.CollectionByName("comments").SimpleFind(&comments, bson.M{"post._id": p.ID}); err != nil {
		return nil
	}
	return comments
}

func (cs *CommentStore) GetByID(id interface{}) *models.Comment {
	var comment models.Comment
	if err := mgm.Coll(&comment).FindByID(id, &comment); err != nil {
		return nil
	}
	return &comment
}

func (cs *CommentStore) Create(c *models.Comment) error {
	return mgm.Coll(c).Create(c)
}

func (cs *CommentStore) Update(c *models.Comment) error {
	return mgm.Coll(c).Update(c)
}

func (cs *CommentStore) Delete(id interface{}) error {
	c := cs.GetByID(id)
	return mgm.Coll(c).Delete(c)
}
