package routers

import (
	"displaylog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/showlogs",&controllers.ShowlogController{},"get:Display")
    beego.Router("/ws",&controllers.WsController{},"get:Wsreturn")
    beego.Router("/specify",&controllers.SpecController{},"get:Specify")
}
