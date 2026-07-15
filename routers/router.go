package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"my-ganji-app/controllers/api"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.CtrlGet("api/user/:id", (*controllers.UserController).GetUserById)

	apiPath := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSRouter("/GetUserInfo", &api.UserController{}, "get:GetUserInfo"),
			beego.NSRouter("/register", &api.UserController{}, "post:UserRegister"),
		))

	beego.AddNamespace(apiPath)
}
