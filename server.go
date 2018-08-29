package main

import (
	configs "checklist/configs"
	controllers "checklist/controllers"
	"fmt"
	"log"
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
		AllowOrigins: []string{"https://tracklist-alpha.surge.sh", "http://tracklist-alpha.surge.sh", "http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	// CORS middleware

	//user accounts
	e.POST("/authenticate", controllers.AuthenticateUser)

	// CList API
	e.POST("/getAllPublicList", controllers.GetAllPublicList)
	e.POST("/search", controllers.SearchListsByText)

	// Restricted
	jwtMiddleWare := middleware.JWT([]byte(jwtConfig.EncryptionKey))
	r := e.Group("/user")
	r.Use(jwtMiddleWare)
	r.POST("/createList", controllers.CreateList)
	r.POST("/getListById", controllers.GetListById)
	r.POST("/deleteList", controllers.DeleteList)
	r.POST("/updateList", controllers.UpdateList)
	r.POST("/getAllList", controllers.GetAllPrivateLists)
	r.POST("/addPublicListToUserList", controllers.AddPublicListToUserList)

	// Tasks API
	e.POST("/getTasksByList", controllers.NotImplemented)
	e.POST("/getTaskById", controllers.NotImplemented)
	e.POST("/addTask", controllers.NotImplemented)
	e.POST("/removeTask", controllers.NotImplemented)
	e.POST("/updateTask", controllers.NotImplemented)

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
