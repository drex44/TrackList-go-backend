package controllers

import (
	configs "checklist/configs"
	tracklistDao "checklist/dao"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var mongoConfig = configs.MongoConfig{}
var jwtConfig = configs.JWTConfig{}
var dao = tracklistDao.TrackListDAO{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	mongoConfig.Read()
	jwtConfig.Read()

	dao.Server = mongoConfig.Server
	dao.Database = mongoConfig.Database
	dao.Connect()
}

//GetUserIDFromJWT ...
func GetUserIDFromJWT(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	return userID
}
