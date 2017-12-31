package main

import "github.com/hal-osaka/akeome/model"

func main() {
	db := model.GetDBConn()

	db.DropTableIfExists(&model.Message{})
	db.AutoMigrate(&model.Message{})
}
