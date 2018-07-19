package controllers

import (
	models "checklist/models"
	modules "checklist/modules"

	"net/http"

	"github.com/labstack/echo"
)

func AuthenticateUser(c echo.Context) error {
	token := new(models.UserSession)
	if err := c.Bind(token); err != nil {
		return err
	}
	isTokenVerified, gResponse := modules.VerifyGoogleTokenID(token.TokenId)
	profile := modules.ConverToUserAccount(gResponse)
	if !isTokenVerified {
		return c.JSON(http.StatusCreated, "{errorCode : 401 , errorMessage : token is not verified}")
	}
	status, id := modules.IsUserExists(profile.Email)
	if !status {
		status, id = modules.CreateNewUser(profile)
		if !status {
			return c.JSON(http.StatusCreated, "{errorCode : 500 , errorMessage : Internal Error. please try again later.}")
		}
	}
	session := modules.CreateSessionToken(id)
	return c.JSON(http.StatusCreated, session)
}
