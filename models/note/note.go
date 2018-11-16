package note

import (
	"demo/models"
	"demo/models/account"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

type images struct {
	Key string `bson:"key" json:"key"`
	Url string `bson:"url" json:"url"`
}
type covers struct {
	Key string `bson:"key" json:"key"`
	Url string `bson:"url" json:"url"`
}

const (
	db         = models.Database
	collection = "note"
)

type Note struct {
	Id             bson.ObjectId   `bson:"_id" json:"id"`
	Title          string          `bson:"title" json:"title"`
	Types          []bson.ObjectId `bson:"types" json:"types"`
	Desc           string          `bson:"desc" json:"desc"`
	ViewCount      uint64          `bson:"viewCount" json:"viewCount"`
	UserShow       bool            `bson:"userShow" json:"userShow"`
	LikedCount     uint64          `bson:"likedCount" json:"likedCount"`
	LikedBy        []bson.ObjectId `bson:"likedBy" json:"likedBy"`
	FollowedCount  uint64          `bson:"followedCount" json:"followedCount"`
	FollowedBy     []bson.ObjectId `bson:"followedBy" json:"followedBy"`
	IsUsedByBrand  bool            `bson:"isUsedByBrand" json:"isUsedByBrand"`
	LinkedProducts []bson.ObjectId `bson:"linkedProducts" json:"linkedProducts"`
	LinkedVendor   bson.ObjectId   `bson:"linkedVendor" json:"linkedVendor"`
	LinkedBrand    bson.ObjectId   `bson:"linkedBrand" json:"linkedBrand"`
	Author         bson.ObjectId   `bson:"author" json:"author"`
	CreatedAt      time.Time       `bson:"createdAt" json:"createdAt"`
	//CreatedBy      *bson.ObjectId   `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
	ModifiedAt time.Time         `bson:"modifiedAt" json:"modifiedAt"`
	ModifiedBy bson.ObjectId     `bson:"modifiedBy" json:"modifiedBy"`
	Status     string            `bson:"status" json:"status"`
	Images     []images          `bson:"images" json:"images"`
	Covers     []covers          `bson:"covers" json:"covers"`
	Xx    []account.Account `bson:"createdBy" json:"createdBy"`
}

func init() {
	index := mgo.Index{
		Key:        []string{"Title"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := models.EnsureIndex(db, collection, index)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func (m *Note) Insert(docs interface{}) error {
	return models.Insert(db, collection, docs)
}

func (m *Note) FindById(id string) (Note, error) {
	var result Note
	//err := models.FindOne(db, collection, bson.M{"_id": bson.ObjectIdHex(id)}, nil, &result)
	pipeline := []bson.M{
		bson.M{"$match": bson.M{"_id": bson.ObjectIdHex(id)}},
		bson.M{"$lookup": bson.M{"from": "accounts", "localField": "createdBy", "foreignField": "_id", "as": "createdBy"}},
	}
	err := models.PipeOne(db, collection, pipeline, &result)
	return result, err
}

func (m *Note) FindAll() ([]Note, error) {
	var result []Note
	pipeline := []bson.M{
		bson.M{"$lookup": bson.M{"from": "accounts", "localField": "createdBy", "foreignField": "_id", "as": "createdBy"}},
	}
	err := models.PipeAll(db, collection, pipeline, &result)
	return result, err
}
