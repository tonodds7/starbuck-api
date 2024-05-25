package userdata

import "fmt"

type IUserData interface {
}

type userDataService struct {
}

func NewUseData() IUserData {
	return &userDataService{}
}

func (service *userDataService) RegisterUser() error {
	fmt.Println("TONNNNNN")
	return nil

}
