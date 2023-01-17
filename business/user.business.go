package business

import (
	"itss.edu.vn/todo_service/core"
	dataaccess "itss.edu.vn/todo_service/data_access"
	"itss.edu.vn/todo_service/models"
)

type UserBusiness struct {
	server *core.Server
	da     *dataaccess.UserDataAccess
}

func NewUserBusiness(server *core.Server) *UserBusiness {
	return &UserBusiness{
		server: server,
		da:     dataaccess.NewUserDataAccess(server),
	}
}

func (b UserBusiness) Get(id string, username string) (*models.User, error) {
	if id != "" {
		return b.da.GetById(id)
	}
	return b.da.GetByUsername(username)
}

func (b UserBusiness) Create(user *models.User) error {
	return b.da.Create(user)
}

func (b UserBusiness) Update(user *models.User) error {
	return b.da.Update(user)
}

func (b UserBusiness) Delete(id string) error {
	return b.da.Delete(id)
}
