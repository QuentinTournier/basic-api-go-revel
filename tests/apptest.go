package tests

import (
	"github.com/revel/revel/testing"
	"io"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) After() {
	println("Tear down")
}

/*
GET /user UserController.Index
POST /user UserController.Create
PUT /user UserController.Update
GET /user/:id UserController.Show
DELETE /user/:id UserController.Delete
 */



func (t *AppTest) TestGetUserRespond(){
	t.Get("/user")
	t.AssertContentType("application/json; charset=utf-8")
	t.AssertOk()
}

func (t *AppTest) TestPostUserRespond(){
	var reader io.Reader
	t.Post("/user","{\"firstName\": \"timeo\",\"lastName\": \"picard\",\"position\": {\"lat\": -40.962026,\"lon\": 24.684415},	\"birthDay\": \"02/24/1994\"}", reader)
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) TestPutUserRespond(){
	var reader io.Reader
	t.Put("/user","{\"firstName\": \"timeo\",\"lastName\": \"picard\",\"position\": {\"lat\": -40.962026,\"lon\": 24.684415},	\"birthDay\": \"02/24/1994\"}", reader)
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) TestDeleteUserRespond(){
	t.Delete("/user")
	t.AssertContentType("text/html; charset=utf-8")
}