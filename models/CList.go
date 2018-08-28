package models

import "gopkg.in/mgo.v2/bson"

type CList struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	Title        string        `bson:"title" json:"title"`
	Description  string        `bson:"description" json:"description"`
	Tags         []string      `bson:"tags" json:"tags"`
	Category     string        `bson:"category" json:"category"`
	Owner        int           `bson:"owner" json:"owner"`
	Contributors []string      `bson:"contributors" json:"contributors"`
	Location     string        `bson:"location" json:"location"`
	Tasks        []Task        `bson:"tasks" json:"tasks"`
	PublicList   bool          `bson:"publiclist" json:"publiclist"`
	Version      int           `bson:"version" json:"version"`
	ParentListID string        `bson:"parentlistid" json:"parentlistid"`
}
type Task struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	Order       int           `bson:"order" json:"order"`
	Status      bool          `bson:"status" json:"status"`
}

type UserLists struct {
	ID    bson.ObjectId   `bson:"_id" json:"id"`
	Tasks []bson.ObjectId `bson:"lists" json:"lists"`
}
