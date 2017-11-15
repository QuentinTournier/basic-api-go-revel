package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/PolytechLyon/cloud-project-equipe-8/app/models/mongodb"
)

type User struct{
ID     bson.ObjectId     `json:"id" bson:"_id"`
Nom       string     `json:"nom" bson:"nom"`
Prenom       string     `json:"prenom" bson:"prenom"`
Age       int     `json:"age" bson:"age"`
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
"nom": m.Nom,"prenom": m.Prenom,"age": m.Age,},

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

// GetUsers Get all User from database and returns
// list of User on success
func GetUsers() ([]User, error) {
	var (
		users []User
		err   error
	)
	c := newUserCollection()
	defer c.Close()
	err = c.Session.Find(nil).All(&users)
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
