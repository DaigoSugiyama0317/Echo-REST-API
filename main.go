package main

import (
	"github.com/DaigoSugiyama0317/Echo-REST-API/controller"
	"github.com/DaigoSugiyama0317/Echo-REST-API/db"
	"github.com/DaigoSugiyama0317/Echo-REST-API/repository"
	"github.com/DaigoSugiyama0317/Echo-REST-API/router"
	"github.com/DaigoSugiyama0317/Echo-REST-API/usecase"
	"github.com/DaigoSugiyama0317/Echo-REST-API/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NesUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}