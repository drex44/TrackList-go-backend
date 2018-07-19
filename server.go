package main

import (
	configs "checklist/configs"
	controllers "checklist/controllers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var jwtConfig = configs.JWTConfig{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	jwtConfig.Read()
}

func main() {
	e := echo.New()

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://tracklist-alpha.herokuapp.com", "https://drex44.github.io"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	// CORS middleware

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//user accounts
	e.POST("/authenticate", controllers.AuthenticateUser)

	// CList API
	e.POST("/getAllPublicList", controllers.GetAllPublicList)
	e.POST("/getListById", controllers.GetListById)
	e.POST("/deleteList", controllers.DeleteList)
	e.POST("/updateList", controllers.UpdateList)
	e.POST("/search", controllers.SearchListsByText)

	// Tasks API
	e.POST("/getTasksByList", controllers.NotImplemented)
	e.POST("/getTaskById", controllers.NotImplemented)
	e.POST("/addTask", controllers.NotImplemented)
	e.POST("/removeTask", controllers.NotImplemented)
	e.POST("/updateTask", controllers.NotImplemented)

	// Restricted
	r := e.Group("/createCList")
	r.Use(middleware.JWT([]byte(jwtConfig.EncryptionKey)))
	r.POST("", controllers.CreateList)

	r = e.Group("/getAllPrivateLists")
	r.Use(middleware.JWT([]byte(jwtConfig.EncryptionKey)))
	r.POST("", controllers.GetAllPrivateLists)

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
