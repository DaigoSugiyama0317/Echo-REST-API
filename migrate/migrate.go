package main

import (
	"fmt"

	"github.com/DaigoSugiyama0317/Echo-REST-API/db"
	"github.com/DaigoSugiyama0317/Echo-REST-API/model"
)

func main() {
	//データベースと接続
	dbConn := db.NewDB()
	defer fmt.Println("successfully Migrated")
	
	//接続の終了
	defer db.CloseDB(dbConn)
	
	//マイグレーションを実行
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}