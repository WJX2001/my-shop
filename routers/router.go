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
			beego.NSRouter("/login", &api.UserController{}, "post:UserLogin"),
			beego.NSRouter("/sendPhoneCode", &api.UserController{}, "post:SendPhoneCode"),
			beego.NSRouter("/phoneCodeCheck", &api.UserController{}, "post:PhoneCodeCheck"),
			beego.NSRouter("/phoneNumberRegisterCheck", &api.UserController{}, "post:PhoneNumberRegisterCheck"),
			beego.NSRouter("/postSendEmailCode", &api.UserController{}, "post:PostSendEmailCode"),
			beego.NSRouter("/emailCodeCheck", &api.UserController{}, "post:EmailCodeCheck"),
		))

	beego.AddNamespace(apiPath)
}
