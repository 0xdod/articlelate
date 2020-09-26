package models

import (
	"github.com/Kamva/mgm/v3"
)

type Favorite struct {
	mgm.DefaultModel `bson:",inline"`
	Article          *Article `json:"article" bson:"article"`
	FavoritedBy      *User    `json:"liked_by" bson:"liked_by"`
}
