package main

import (
	. "checklist/configs"
	. "checklist/dao"
	. "checklist/models"
	"net/http"

	"github.com/labstack/echo"
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

	// CList API
	e.POST("/createCList", createCList)
	e.POST("/getAllCList", getAllCList)
	e.POST("/getCListById", getCListById)
	e.POST("/deleteCList", deleteCList)
	e.POST("/updateCList", updateCList)

	// Items API
	e.POST("/getItemsByCList", notImplemented)
	e.POST("/getItemById", notImplemented)
	e.POST("/addItem", notImplemented)
	e.POST("/removeItem", notImplemented)
	e.POST("/updateItem", notImplemented)

	e.Logger.Fatal(e.Start(":4000"))
}

func createCList(c echo.Context) error {
	u := new(CList)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := dao.Insert(*u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func deleteCList(c echo.Context) error {
	u := new(CList)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := dao.Delete(*u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func getCListById(c echo.Context) error {
	u := new(CList)
	if err := c.Bind(u); err != nil {
		return err
	}
	list, err := dao.FindById(*u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, list)
}

func getAllCList(c echo.Context) error {
	var lists, err = dao.FindAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, lists)
}

func updateCList(c echo.Context) error {
	u := new(CList)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := dao.Update(*u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func notImplemented(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
