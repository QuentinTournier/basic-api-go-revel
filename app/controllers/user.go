package controllers

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"github.com/revel/revel"
	//"encoding/json"
	"github.com/PolytechLyon/cloud-project-equipe-8/app/models"
	"strconv"
	"github.com/kpawlik/geojson"
)

type UserController struct{
 *revel.Controller
}

func (c UserController) Index() revel.Result {
	var (
		users []models.User
		err error
		page int
	)

	page, err = strconv.Atoi(c.Params.Get("page"))

	if err != nil {
		page = 0
	} else if page < 0 {
		page = 0
	}

	users, err = models.GetUsers(page)
	if err != nil{
		errResp := buildErrResponse(err,"500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200

	if(users == nil){
		users = []models.User{};
	}

    return c.RenderJSON(&users)
}

func (c UserController) FindByName() revel.Result {
	var (
		users []models.User
		err error
		page int
		name string
	)

	page, err = strconv.Atoi(c.Params.Get("page"))
	if err != nil {
		page = 0
	} else if page < 0 {
		page = 0
	}

	name = c.Params.Get("term")

	users, err = models.GetUsersByName(name, page)
	if err != nil{
		errResp := buildErrResponse(err,"500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	if(users == nil){
		users = []models.User{};
	}

	c.Response.Status = 200
	return c.RenderJSON(&users)
}

func (c UserController) FindByAge() revel.Result {
	var (
		users []models.User
		err error
		page int
		age int
	)

	page, err = strconv.Atoi(c.Params.Get("page"))
	if err != nil {
		page = 0
	} else if page < 0 {
		page = 0
	}

	age, err = strconv.Atoi(c.Params.Get("eq"))
	if err == nil {
		users, err = models.GetUsersByAgeEq(age, page)
	} else {
		age, err = strconv.Atoi(c.Params.Get("gt"))
		if err == nil {
			users, err = models.GetUsersByAgeGt(age, page)
		} else {
			err = errors.New("Invalid age args")
		}
	}

	if err != nil{
		errResp := buildErrResponse(err,"500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	if(users == nil){
		users = []models.User{};
	}

	c.Response.Status = 200
	return c.RenderJSON(&users)
}

func (c UserController) FindByPosition() revel.Result {
	var (
		users []models.User
		err error
		page int
		lon float64
		lat float64
	)

	page, err = strconv.Atoi(c.Params.Get("page"))
	if err != nil {
		page = 0
	} else if page < 0 {
		page = 0
	}

	lon, err = strconv.ParseFloat(c.Params.Get("lon"), 64)
	if err != nil {
		lon = 0
	}

	lat, err = strconv.ParseFloat(c.Params.Get("lat"), 64)
	if err != nil {
		lat = 0
	}

	users, err = models.GetUsersByPosition(geojson.Coordinate{geojson.CoordType(lon), geojson.CoordType(lat)}, page)
	if err != nil{
		errResp := buildErrResponse(err,"500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	if(users == nil){
		users = []models.User{};
	}

	c.Response.Status = 200
	return c.RenderJSON(&users)
}

func (c UserController) Show(id string) revel.Result {
    var (
    	user models.User
    	err error
    	userID bson.ObjectId
    )

    if id == "" {
    	errResp := buildErrResponse(errors.New("Invalid user id format 'null'"),"400")
    	c.Response.Status = 400
    	return c.RenderJSON(errResp)
    }

    userID, err = convertToObjectIdHex(id)
    if err != nil {
	    println(err)
    	errResp := buildErrResponse(errors.New("Invalid user id format"),"404")
    	c.Response.Status = 404
    	return c.RenderJSON(errResp)
    }

    user, err = models.GetUser(userID)
    if err != nil{
    	errResp := buildErrResponse(err,"404")
    	c.Response.Status = 404
    	return c.RenderJSON(errResp)
    }

    c.Response.Status = 200
    return c.RenderJSON(&user)
}

func (c UserController) Create(user *models.User) revel.Result {
	    var (
		err error
	    	//createdUser models.User
	    )

	*user, err = models.AddUser(*user)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}

	//user.ID = createdUser.ID

    c.Response.Status = 201
    return c.RenderJSON(&user)
}

func (c UserController) Update(id string, user *models.User) revel.Result {
	var (
		//user models.User
		userID bson.ObjectId
		err error
    	)

	userID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid user id format"),"400")
		c.Response.Status = 400
		return c.RenderJSON(errResp)
	}

	user.ID = userID
	err = user.UpdateUser()
	if err != nil{
		errResp := buildErrResponse(err,"500")
    		c.Response.Status = 500
    		return c.RenderJSON(errResp)
	}

    	return c.RenderJSON(&user)
}

func (c UserController) Delete(id string) revel.Result {
	var (
    	err error
    	user models.User
    	userID bson.ObjectId
    )
     if id == ""{
    	errResp := buildErrResponse(errors.New("Invalid user id format 'null'"),"400")
    	c.Response.Status = 400
    	return c.RenderJSON(errResp)
    }

    userID, err = convertToObjectIdHex(id)
    if err != nil{
    	errResp := buildErrResponse(errors.New("Invalid user id format"),"500")
    	c.Response.Status = 500
    	return c.RenderJSON(errResp)
    }

    user, err = models.GetUser(userID)
    if err != nil{
    	errResp := buildErrResponse(err,"500")
    	c.Response.Status = 500
    	return c.RenderJSON(errResp)
    }
	err = user.DeleteUser()
	if err != nil{
		errResp := buildErrResponse(err,"500")
    		c.Response.Status = 500
    		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200

    return c.RenderJSON(Empty{})
}

func (c UserController) DeleteAll() revel.Result {
	var (
    	err error
    )

	err = models.DeleteAllUser()
	if err != nil{
		errResp := buildErrResponse(err,"500")
		c.Response.Status = 500
		return c.RenderJSON(errResp)
	}
	c.Response.Status = 200

    return c.RenderJSON(Empty{})
}


func (c UserController) CreateAll() revel.Result {
	var (
		err error
		users []*models.User
	)

	c.Params.BindJSON(&users)

	for _,user := range users {
		*user, err = models.AddUser(*user)

		if err != nil{
			errResp := buildErrResponse(err,"500")
			c.Response.Status = 500
			return c.RenderJSON(errResp)
		}
	}

	if(users == nil){
		users = []*models.User{};
	}

	c.Response.Status = 201
	return c.RenderJSON(&users)
}

type Empty struct {
}