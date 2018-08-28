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
	PublicListCollection = "PublicList"
	UserCollection       = "User"
)

// Connect ...
//connect to the db
func (m *TrackListDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// lists collection

//FindAllPublicLists ...
func (m *TrackListDAO) FindAllPublicLists() ([]CList, error) {
	var clists []CList
	err := db.C(PublicListCollection).Find(bson.M{}).All(&clists)
	return clists, err
}

//FindAllPrivateLists ...
func (m *TrackListDAO) FindAllPrivateLists(userID string) ([]CList, error) {
	var profile UserAccount
	err := db.C(UserCollection).FindId(bson.ObjectIdHex(userID)).One(&profile)
	if err != nil {
		fmt.Println(err)
	}
	return profile.Lists, err
}

//FindListByID ...
func (m *TrackListDAO) FindListByID(userID string, listID string) (CList, error) {

	// match := bson.M{"$match": bson.M{"_id": bson.ObjectIdHex(userID), "lists._id": bson.ObjectIdHex(listID)}}

	// condition := bson.M{"$eq": []string{"$$list._id", "ObjectId('" + listID + "')"}}
	// filter := bson.M{"input": "$lists", "as": "list", "cond": condition}
	// project := bson.M{"$project": bson.M{"lists": bson.M{"$filter": filter}}}
	//pipe := db.C(UserCollection).Pipe([]bson.M{match, project})
	var list CList
	lists, err := m.FindAllPrivateLists(userID)

	if err != nil {
		fmt.Println(err)
		return list, err
	}
	for i := 0; i < len(lists); i++ {
		if lists[i].ID == bson.ObjectIdHex(listID) {
			list = lists[i]
			break
		}
	}
	return list, err
}

//FindPublicListByID ...
func (m *TrackListDAO) FindPublicListByID(listID string) (CList, error) {

	var list CList
	err := db.C(PublicListCollection).Find(bson.M{"_id": bson.ObjectIdHex(listID)}).One(&list)
	return list, err
}

//AddPublicListToUserList ...
func (m *TrackListDAO) AddPublicListToUserList(userID string, listID string) (CList, error) {

	list, err := m.FindPublicListByID(listID)
	fmt.Println(list)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	list.ParentListID = list.ID.Hex()
	list, err = m.InsertNewList(userID, list)
	if err != nil {
		fmt.Println(err)
		return list, err
	}

	return list, err
}

//InsertNewList ...
func (m *TrackListDAO) InsertNewList(userID string, clist CList) (CList, error) {
	listID := bson.NewObjectId()
	clist.ID = listID
	for index := 0; index < len(clist.Tasks); index++ {
		clist.Tasks[index].ID = bson.NewObjectId()
	}
	clist.Version = 0

	who := bson.M{"_id": bson.ObjectIdHex(userID)}
	PushToArray := bson.M{"$push": bson.M{"lists": clist}}

	errList := db.C(UserCollection).Update(who, PushToArray)

	if errList != nil {
		return clist, errList
	}

	return clist, errList
}

//DeleteList ...
func (m *TrackListDAO) DeleteList(userID string, listID bson.ObjectId) error {

	who := bson.M{"_id": bson.ObjectIdHex(userID)}
	change := bson.M{"$pull": bson.M{"lists": bson.M{"_id": listID}}}

	err := db.C(UserCollection).Update(who, change)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

//UpdateList ...
func (m *TrackListDAO) UpdateList(userID string, clist CList) (CList, error) {

	for index := 0; index < len(clist.Tasks); index++ {
		if clist.Tasks[index].ID == "" {
			clist.Tasks[index].ID = bson.NewObjectId()
		}
	}
	clist.Version++

	if clist.PublicList {
		clist.PublicList = false
		err := db.C(PublicListCollection).Insert(&clist)
		if err != nil {
			if mgo.IsDup(err) {
				err := db.C(PublicListCollection).Update(bson.M{"_id": clist.ID}, clist)
				if err != nil {
					fmt.Println(err)
					return clist, err
				}
			} else {
				fmt.Println(err)
				return clist, err
			}
		}
		clist.PublicList = true
	}

	who := bson.M{"_id": bson.ObjectIdHex(userID), "lists._id": clist.ID}
	change := bson.M{"$set": bson.M{"lists.$": clist}}
	err := db.C(UserCollection).Update(who, change)

	if err != nil {
		fmt.Println(err)
		return clist, err
	}
	return clist, err
}

//SearchLists ...
func (m *TrackListDAO) SearchLists(text string) ([]CList, error) {
	var lists []CList
	err := db.C(PublicListCollection).Find(bson.M{"$text": bson.M{"$search": text}}).All(&lists)
	return lists, err
}

// lists collection

//users collections

//FindUserByEmail ...
//find user by email address
func (m *TrackListDAO) FindUserByEmail(email string) (UserAccount, error) {
	var profile UserAccount
	err := db.C(UserCollection).Find(bson.M{"email": email}).One(&profile)
	return profile, err
}

//InsertNewUser ...
func (m *TrackListDAO) InsertNewUser(profile UserAccount) (UserAccount, error) {
	var id = bson.NewObjectId()
	profile.ID = id
	err := db.C(UserCollection).Insert(&profile)
	return profile, err
}
