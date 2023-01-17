package main

import (
	"itss.edu.vn/todo_service/apis"
	"itss.edu.vn/todo_service/core"
	"itss.edu.vn/todo_service/utils"
)

func main() {
	server := core.NewServer()

	server.Echo.Validator = utils.NewValidator()

	_ = apis.NewHealthyAPI("/healthy", server)

	_ = apis.NewTaskApi("/tasks", server)

	_ = apis.NewUserApi("/users", server)

	server.Start()
}
