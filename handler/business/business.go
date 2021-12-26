package business

import (
	"github.com/sachinmahanin/passwordRepeated/config"
	"github.com/sachinmahanin/passwordRepeated/handler/business/model"

	webserver "github.com/zhongjie-cai/web-server"
)

var MAX_QUEUE_SIZE = 10

func Lookup(session webserver.Session) (interface{}, error) {

	session.LogMethodLogic(
		webserver.LogLevelInfo,
		"business",
		"Lookup", "Lookup",
	)

	var passwordRequest model.PasswordRequest
	var bodyError = session.GetRequestBody(
		&passwordRequest,
	)
	if bodyError != nil {
		return nil, bodyError
	}
	//if the string exist
	for temp := config.PasswordHistory.Front(); temp != nil; temp = temp.Next() {
		if temp.Value == passwordRequest.Password {
			return "password used recently", nil
		}
	}

	if config.PasswordHistory.Len() == MAX_QUEUE_SIZE {
		config.PasswordHistory.Remove(config.PasswordHistory.Front())
	}
	config.PasswordHistory.PushBack(passwordRequest.Password)

	return "password not recently used", nil
}
