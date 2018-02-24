package postctl

import "github.com/sj14/web-demo/interfaces/web/controller/mainctl"

func NewPostController(utilController mainctl.MainController) PostController {
	return PostController{utilController}
}

type PostController struct {
	mainctl.MainController
}
