package handler

import (
	"backend/pkg/userdata"
	"net/http"
)

type IUseDataHandler struct{}

type UseDataHandler struct {
	UserDataService userdata.IUserData
}

func NewUseDataHandler() *UseDataHandler {
	return &UseDataHandler{
		userdata.NewUserData(),
	}
}

func (handler *UseDataHandler) UserDatasController(req *http.Request) {
	handler.UserDataService.RegisterUser()

}
