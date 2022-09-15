package beego

import (
	"github.com/beego/beego/v2/server/web"
	"testing"
)

func TestMainController_PutSession(t *testing.T) {
	webConfig()

	ctrl := &MainController{}
	web.Router("/session", ctrl, "post:PutSession")
}

func webConfig() {
	web.BConfig.WebConfig.Session.SessionOn = true
	web.BConfig.WebConfig.Session.SessionGCMaxLifetime = 10
}
