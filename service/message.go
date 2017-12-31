package service

import "github.com/alberii/akeome/model"

type message struct {
}

var Message = message{}

func (self *message) Store(message model.Message) model.Message {
	db.Create(&message)
	return message
}
func (self *message) FindByTwitter(tid string) []model.Message {
	var message []model.Message
	db.Where("twitter_id = ?", tid).Find(&message)
	return message
}
