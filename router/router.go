package router

import (
	"github.com/DaigoSugiyama0317/Echo-REST-API/controller"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUP)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	return e
}