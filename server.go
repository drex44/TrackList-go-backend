package main

import (
	"net/http"
	"github.com/labstack/echo"
	. "checklist/models"
	. "checklist/dao"
	. "checklist/configs"
)

var mongoConfig = MongoConfig{}
var dao = ListDAO{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	mongoConfig.Read()

	dao.Server = mongoConfig.Server
	dao.Database = mongoConfig.Database
	dao.Connect()
}

func main() {
	e := echo.New()
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	e.POST("/createCList", createCList)
	e.GET("/updateCList", updateCList)
	e.GET("/addItemToCList", addItemToCList)
	e.GET("/getAll", getAll)

	e.Logger.Fatal(e.Start(":1323"))
}

func createCList(c echo.Context) error {
	u := new(CList)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err:=dao.Insert(*u); err!=nil{
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func getAll(c echo.Context) error {
	var lists, err=dao.FindAll()
	if err != nil {
		return err
	}
  	return c.JSON(http.StatusOK, lists)
}

func updateCList(c echo.Context) error {
  	return c.String(http.StatusOK, "Hello, World!")
}

func addItemToCList(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}