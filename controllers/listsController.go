package controllers

import (
	models "checklist/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func CreateList(c echo.Context) error {

	userID := GetUserIDFromJWT(c)

	u := new(models.CList)
	if err := c.Bind(u); err != nil {
		return err
	}
	clist, err := dao.InsertNewList(userID, *u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, clist)
}

func DeleteList(c echo.Context) error {
	u := new(models.CList)
	userID := GetUserIDFromJWT(c)

	if err := c.Bind(u); err != nil {
		return err
	}
	if err := dao.DeleteList(userID, u.ID); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func GetListById(c echo.Context) error {

	userID := GetUserIDFromJWT(c)
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		//json_map has the JSON Payload decoded into a map
		listID := json_map["id"].(string)
		list, err := dao.FindListByID(userID, listID)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return c.JSON(http.StatusCreated, list)
	}

}

func AddPublicListToUserList(c echo.Context) error {

	userID := GetUserIDFromJWT(c)
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		//json_map has the JSON Payload decoded into a map
		listID := json_map["listid"].(string)
		list, err := dao.AddPublicListToUserList(userID, listID)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return c.JSON(http.StatusCreated, list)
	}

}

func GetAllPublicList(c echo.Context) error {
	var lists, err = dao.FindAllPublicLists()
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
	userID := GetUserIDFromJWT(c)
	u := new(models.CList)
	if err := c.Bind(u); err != nil {
		return err
	}
	list, err := dao.UpdateList(userID, *u)
	if err != nil {
		return c.JSON(http.StatusCreated, err)
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
