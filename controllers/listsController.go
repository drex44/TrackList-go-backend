package controllers

import (
	models "checklist/models"
	"net/http"

	"github.com/labstack/echo"
)

func CreateList(c echo.Context) error {

	userID := GetUserIDFromJWT(c)

	u := new(models.CList)
	if err := c.Bind(u); err != nil {
		return err
	}
	clist, err := dao.InsertNewList(*u, userID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, clist)
}

func DeleteList(c echo.Context) error {
	u := new(models.CList)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := dao.DeleteList(*u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func GetListById(c echo.Context) error {
	u := new(models.CList)
	if err := c.Bind(u); err != nil {
		return err
	}
	list, err := dao.FindListByID(*u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, list)
}

func GetAllPublicList(c echo.Context) error {
	var lists, err = dao.FindAllLists()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, lists)
}

func GetAllPrivateLists(c echo.Context) error {
	userID := GetUserIDFromJWT(c)
	var lists, err = dao.FindAllPrivateLists(userID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, lists)
}

func UpdateList(c echo.Context) error {
	u := new(models.CList)
	if err := c.Bind(u); err != nil {
		return err
	}
	list, err := dao.UpdateList(*u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, list)
}

func SearchListsByText(c echo.Context) error {
	u := new(models.SearchText)
	if err := c.Bind(u); err != nil {
		return err
	}
	lists, err := dao.SearchLists(u.Text)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, lists)
}

func NotImplemented(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
