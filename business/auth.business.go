package business

import (
	"context"

	"firebase.google.com/go/v4/auth"
	"itss.edu.vn/todo_service/core"
	"itss.edu.vn/todo_service/models"
)

type AuthBusiness struct {
	server *core.Server
}

func NewAuthBusiness(s *core.Server) *AuthBusiness {
	return &AuthBusiness{
		server: s,
	}
}

func (b AuthBusiness) Register(regUser *models.UserRegistrationRequest) error {
	newUser := &auth.UserToCreate{}
	newUser.Email(regUser.Email)
	newUser.Password(regUser.Password)
	_, err := b.server.AuthClient.CreateUser(context.Background(), newUser)

	return err
}
