package api

import (
	"fmt"
	"myproject/configs"
	"myproject/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

var user []models.User

func Index(c echo.Context) error {
	rows := configs.CreateCon().Find(&user)
	err := configs.RedisCon().Set("users_index", rows, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	user_index, err := configs.RedisCon().Get("users_index").Result()
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, user_index)
}

func Show(c echo.Context) error {
	id := c.Param("id")
	prod := configs.CreateCon().Find(&models.User{}, id)
	return c.JSON(http.StatusOK, prod)
}

func Store(c echo.Context) error {
	first_name := c.FormValue("first_name")
	full_name := c.FormValue("full_name")
	email := c.FormValue("email")
	mobile := c.FormValue("mobile")
	status, _ := strconv.ParseInt(c.FormValue("status"), 10, 64)
	password := c.FormValue("psassword")

	user := configs.CreateCon().Create(&models.User{
		FirstName: first_name,
		LastName:  full_name,
		Email:     &email,
		Mobile:    mobile,
		VerifyAt:  time.Now(),
		Status:    uint(status),
		Password:  password,
	})

	return c.JSON(http.StatusOK, user)
}

/*func Update(c echo.Context) error {

	user_id := c.Param("id")

	first_name := c.FormValue("first_name")
	last_name := c.FormValue("full_name")
	email := c.FormValue("email")
	mobile := c.FormValue("mobile")
	status, _ := strconv.ParseInt(c.FormValue("status"), 10, 64)
	password := c.FormValue("password")
	person := configs.CreateCon().First(&models.User{}, user_id).Update(models.User{
		FirstName: first_name,
		LastName:  last_name,
		Email:     &email,
		Mobile:    mobile,
	     VerifyAt:  time.Now(),
		Status:    uint(status),
		Password:  password,
	})

	return c.JSON(http.StatusOK, person)
}*/

func Delete(c echo.Context) error {

	user_id := c.Param("id")

	person := configs.CreateCon().Delete(&models.User{}, user_id)

	return c.JSON(http.StatusOK, person)
}
