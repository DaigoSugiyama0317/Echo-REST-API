package main

import (
	"github.com/DaigoSugiyama0317/Echo-REST-API/controller"
	"github.com/DaigoSugiyama0317/Echo-REST-API/db"
	"github.com/DaigoSugiyama0317/Echo-REST-API/repository"
	"github.com/DaigoSugiyama0317/Echo-REST-API/router"
	"github.com/DaigoSugiyama0317/Echo-REST-API/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	useController := controller.NesUserController(userUsecase)
	e := router.NewRouter(useController)
	e.Logger.Fatal(e.Start(":8080"))
}