package dao

import (
	. "checklist/models"
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ListDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "lists"
)

func (m *ListDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *ListDAO) FindAll() ([]CList, error) {
	var clists []CList
	err := db.C(COLLECTION).Find(bson.M{}).All(&clists)
	return clists, err
}

func (m *ListDAO) FindById(clist CList) (CList, error) {
	var resCList CList
	err := db.C(COLLECTION).FindId(clist.ID).One(&resCList)
	return resCList, err
}

func (m *ListDAO) Insert(clist CList) error {
	clist.ID = bson.NewObjectId()
	for index := 0; index < len(clist.Tasks); index++ {
		clist.Tasks[index].ID = bson.NewObjectId()
	}
	fmt.Println(clist)
	err := db.C(COLLECTION).Insert(&clist)
	fmt.Println(err)
	return err
}

func (m *ListDAO) Delete(clist CList) error {
	fmt.Println(clist)
	err := db.C(COLLECTION).RemoveId(clist.ID)
	fmt.Println(err)
	return err
}

func (m *ListDAO) Update(clist CList) error {
	err := db.C(COLLECTION).UpdateId(clist.ID, &clist)
	return err
}
