package main

import (
	. "checklist/configs"
	. "checklist/dao"
	. "checklist/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://tracklist-alpha.herokuapp.com", "https://drex44.github.io"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// CORS middleware

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// CList API
	e.POST("/createCList", createCList)
	e.POST("/getAllCList", getAllCList)
	e.POST("/getCListById", getCListById)
	e.POST("/deleteCList", deleteCList)
	e.POST("/updateCList", updateCList)
	e.POST("/search", searchByText)

	// Tasks API
	e.POST("/getTasksByCList", notImplemented)
	e.POST("/getTaskById", notImplemented)
	e.POST("/addTask", notImplemented)
	e.POST("/removeTask", notImplemented)
	e.POST("/updateTask", notImplemented)

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	e.Logger.Fatal(e.Start(addr))
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func createCList(c echo.Context) error {
	u := new(CList)
	if err := c.Bind(u); err != nil {
		return err
	}
	clist, err := dao.Insert(*u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, clist)
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
	list, err := dao.Update(*u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, list)
}

func notImplemented(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func searchByText(c echo.Context) error {
	u := new(SearchText)
	if err := c.Bind(u); err != nil {
		return err
	}
	fmt.Println(u.Text)
	lists, err := dao.Search(u.Text)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, lists)
}
