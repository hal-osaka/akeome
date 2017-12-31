package service

import (
	"testing"

	"github.com/hal-osaka/akeome/model"
)

func TestMessage_Store(t *testing.T) {
	Message.Store(model.Message{
		TwitterId: "makki0205",
		Body:      "hoge--",
	})
}
