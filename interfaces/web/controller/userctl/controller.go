package userctl

import "github.com/sj14/web-demo/interfaces/web/controller/mainctl"

func NewUserController(utilController mainctl.MainController) UserController {
	return UserController{utilController}
}

type UserController struct {
	mainctl.MainController
}
