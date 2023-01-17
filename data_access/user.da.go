package dataaccess

import (
	"itss.edu.vn/todo_service/core"
	"itss.edu.vn/todo_service/models"
)

type UserDataAccess struct {
	server *core.Server
}

func NewUserDataAccess(server *core.Server) *UserDataAccess {
	return &UserDataAccess{
		server: server,
	}
}

func (d UserDataAccess) GetById(id string) (*models.User, error) {
	user := &models.User{}
	if err := d.server.Db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d UserDataAccess) GetByUsername(username string) (*models.User, error) {
	user := &models.User{}
	if err := d.server.Db.Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d UserDataAccess) Create(user *models.User) error {
	return d.server.Db.Create(user).Error
}

func (u UserDataAccess) Update(user *models.User) error {
	return u.server.Db.Save(user).Error
}

func (u UserDataAccess) Delete(id string) error {
	return u.server.Db.Delete(&models.User{}, map[string]interface{}{
		"id": id,
	}).Error
}
