package main

import (
	"fmt"

	"github.com/DaigoSugiyama0317/Echo-REST-API/db"
	"github.com/DaigoSugiyama0317/Echo-REST-API/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}