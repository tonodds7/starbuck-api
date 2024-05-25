package userdata

import "fmt"

type IUserData interface {
	RegisterUser() error
}

type userDataService struct {
}

func NewUserData() IUserData {
	return &userDataService{}
}

func (service *userDataService) RegisterUser() error {
	fmt.Println("TONNNNNN")
	return nil

}
