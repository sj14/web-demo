package profilectl

import "github.com/sj14/web-demo/interfaces/web/controller/mainctl"

func NewProfileController(utilController mainctl.MainController) ProfileController {
	return ProfileController{utilController}
}

type ProfileController struct {
	mainctl.MainController
}
