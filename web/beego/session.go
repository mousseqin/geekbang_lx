package beego

import "github.com/beego/beego/v2/server/web"

type MainController struct {
	web.Controller
}

func (ctrl *MainController) PutSession() {
	_ = ctrl.SetSession("name", "web session")

	ctrl.TplName = "hello_world.html"
	ctrl.Data["name"] = "PutSession"
	_ = ctrl.Render()
}
