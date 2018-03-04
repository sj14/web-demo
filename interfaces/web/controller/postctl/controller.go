package postctl

import "github.com/sj14/web-demo/interfaces/web/controller/mainctl"

func NewPostController(mainController mainctl.MainController) PostController {
	return PostController{mainController}
}

type PostController struct {
	mainctl.MainController
}
