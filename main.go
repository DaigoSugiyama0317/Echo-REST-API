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
	// DB接続の初期化
	db := db.NewDB()

	// バリデーションのインスタンス生成
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	// リポジトリ層の初期化
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	// ユースケース（ビジネスロジック）層
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)

	// コントローラー層の初期化
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	
	// ルーターを構築して、エンドポイントを登録
	e := router.NewRouter(userController, taskController)

	// サーバー起動（:8080で待ち受け）
	e.Logger.Fatal(e.Start(":8080"))
}
