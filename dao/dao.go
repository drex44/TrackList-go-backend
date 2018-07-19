package dao

import (
	. "checklist/models"
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//TrackListDAO ...
// dao object
type TrackListDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// lists of collections
const (
	ListsCollection     = "lists"
	UsersCollection     = "users"
	UserListsCollection = "UserLists"
)

func (m *TrackListDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// lists collection

//FindAllLists ...
func (m *TrackListDAO) FindAllLists() ([]CList, error) {
	var clists []CList
	err := db.C(ListsCollection).Find(bson.M{}).All(&clists)
	return clists, err
}

//FindAllPrivateLists ...
func (m *TrackListDAO) FindAllPrivateLists(userID string) ([]CList, error) {
	var clists []CList
	userLists, err := FindUserLists(userID)
	if err != nil {
		fmt.Println(err)
		return clists, err
	}

	err = db.C(ListsCollection).Find(bson.M{"_id": bson.M{"$in": userLists.Tasks}}).All(&clists)
	if err != nil {
		fmt.Println(err)
	}
	return clists, err
}

//FindListByID ...
func (m *TrackListDAO) FindListByID(clist CList) (CList, error) {
	var resCList CList
	err := db.C(ListsCollection).FindId(clist.ID).One(&resCList)
	return resCList, err
}

//InsertNewList ...
func (m *TrackListDAO) InsertNewList(clist CList, userID string) (CList, error) {
	listID := bson.NewObjectId()
	clist.ID = listID
	for index := 0; index < len(clist.Tasks); index++ {
		clist.Tasks[index].ID = bson.NewObjectId()
	}
	errList := db.C(ListsCollection).Insert(&clist)

	if errList == nil {
		userLists, err := FindUserLists(userID)
		if err != nil {
			userLists.ID = bson.ObjectIdHex(userID)
		}
		userLists.Tasks = append(userLists.Tasks, listID)

		err = db.C(UserListsCollection).Insert(&userLists)
		if err != nil {
			fmt.Println(err)
			return clist, err
		}
	}

	return clist, errList
}

//DeleteList ...
func (m *TrackListDAO) DeleteList(clist CList) error {
	err := db.C(ListsCollection).RemoveId(clist.ID)
	return err
}

//UpdateList ...
func (m *TrackListDAO) UpdateList(clist CList) (CList, error) {

	for index := 0; index < len(clist.Tasks); index++ {
		if clist.Tasks[index].ID == "" {
			clist.Tasks[index].ID = bson.NewObjectId()
		}
	}

	err := db.C(ListsCollection).UpdateId(clist.ID, &clist)
	return clist, err
}

//SearchLists ...
func (m *TrackListDAO) SearchLists(text string) ([]CList, error) {
	var lists []CList
	err := db.C(ListsCollection).Find(bson.M{"$text": bson.M{"$search": text}}).All(&lists)
	return lists, err
}

// lists collection

//users collections

//FindUserByEmail ...
//find user by email address
func (m *TrackListDAO) FindUserByEmail(email string) (UserAccount, error) {
	var profile UserAccount
	err := db.C(UsersCollection).Find(bson.M{"email": email}).One(&profile)
	return profile, err
}

//InsertNewUser ...
func (m *TrackListDAO) InsertNewUser(profile UserAccount) (UserAccount, error) {
	var id = bson.NewObjectId()
	profile.ID = id
	err := db.C(UsersCollection).Insert(&profile)
	return profile, err
}

//users collections

// lists of user collection

//FindUserLists ...
func FindUserLists(ID string) (UserLists, error) {
	var res UserLists
	err := db.C(UserListsCollection).FindId(bson.ObjectIdHex(ID)).One(&res)
	if err != nil {
		fmt.Println(err)
	}
	return res, err
}

// lists of user collection
