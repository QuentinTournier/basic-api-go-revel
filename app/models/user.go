package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/PolytechLyon/cloud-project-equipe-8/app/models/mongodb"
	"time"
	"encoding/json"
	"github.com/kpawlik/geojson"
)

const ctLayoutAccept = "1/2/2006"
const ctLayoutReturn = "01/02/2006"

type User struct{
	ID     bson.ObjectId     `json:"id" bson:"_id"`
	FirstName       string     `json:"firstName" bson:"firstName"`
	LastName       string     `json:"lastName" bson:"lastName"`
	BirthDay       time.Time     `json:"birthDay" bson:"birthDay"`
	Position       JSONPoint     `json:"position" bson:"position"`
}


func newUserCollection()  *mongodb.Collection  {
   return mongodb.NewCollectionSession("users")
}

// AddUser insert a new User into database and returns
// last inserted user on success.
func AddUser(m User) (user User, err error) {
	c := newUserCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	return m, c.Session.Insert(m)
}

// UpdateUser update a User into database and returns
// last nil on success.
func (m User) UpdateUser() error{
	c := newUserCollection()
	defer c.Close()
	
	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
"firstName": m.FirstName,"lastName": m.LastName,"birthDay": m.BirthDay,"position": m.Position,},

	})
	return err
}

// DeleteUser Delete User from database and returns
// last nil on success.
func (m User) DeleteUser() error{
	c := newUserCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// DeleteUser Delete User from database and returns
// last nil on success.
func DeleteAllUser() error{
	c := newUserCollection()
	defer c.Close()

	_, err := c.Session.RemoveAll(bson.D{})
	return err
}

// GetUsers Get all User from database and returns
// list of User on success
func GetUsers(page int) ([]User, error) {
	return GetUsersWithQuery(nil, page)
}

func GetUsersByName(name string, page int) ([]User, error) {
	return GetUsersWithQuery(bson.M{"lastName": bson.M{"$regex": name}}, page)
}

func GetUsersByPosition(position geojson.Coordinate, page int) ([]User, error) {
	return GetUsersWithQuery(bson.M{"position": bson.M{"coordinates": bson.M{"$near": position}}}, page)
}

func GetUsersByDate(date time.Time, selector string, page int) ([]User, error) {
	return GetUsersWithQuery(bson.M{"birthDay": bson.M{selector: date}}, page)
}

func GetUsersByAge(age int, selector string, page int) ([]User, error) {
	now := time.Now()
	then := now.AddDate(-age, 0, 0)

	return GetUsersByDate(then, selector, page)
}

func GetUsersByAgeEq(age int, page int) ([]User, error) {
	return GetUsersByAge(age, "$eq", page)
}

func GetUsersByAgeGt(age int, page int) ([]User, error) {
	return GetUsersByAge(age, "$lt", page)
}

func GetUsersByAgeLt(age int, page int) ([]User, error) {
	return GetUsersByAge(age, "$gt", page)
}

func GetUsersWithQuery(query interface{}, page int) ([]User, error) {
	var (
		users []User
		err   error
	)

	c := newUserCollection()
	defer c.Close()

	err = c.Session.Find(query).Sort("-birthDay").Skip(page*100).Limit(100).All(&users)
	return users, err
}

// GetUser Get a User from database and returns
// a User on success
func GetUser(id bson.ObjectId) (User, error) {
	var (
		user User
		err   error
	)

	c := newUserCollection()
	defer c.Close()


	err = c.Session.Find(bson.M{"_id": id}).One(&user)
	return user, err
}

func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User

	return json.Marshal(&struct {
		BirthDay string `json:"birthDay"`
		*Alias
	}{
		BirthDay: u.BirthDay.Format(ctLayoutReturn),
		Alias:    (*Alias)(u),
	})
}

func (u *User) UnmarshalJSON(data []byte) error {
	type Alias User
	aux := &struct {
		BirthDay string `json:"birthDay"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var err error
	u.BirthDay, err = time.Parse(ctLayoutAccept, aux.BirthDay)
	return err
}