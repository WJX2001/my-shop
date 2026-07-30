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
			beego.NSRouter("/postEmailCheck", &api.UserController{}, "post:PostEmailCheck"),
			beego.NSRouter("/bindFundPassword", &api.UserController{}, "post:BindFundPassword"),
			beego.NSRouter("/updatePassword", &api.UserController{}, "post:UpdatePassword"),
			beego.NSRouter("/updateCreatePhoneEmail", &api.UserController{}, "post:UpdateCreatePhoneEmail"),
			beego.NSRouter("/forgetPassword", &api.UserController{}, "post:ForgetPassword"),
		),
		beego.NSNamespace("/userInfo",
			beego.NSRouter("/getUserInfo", &api.UserInfoController{}, "post:GetUserInfo"),
			beego.NSRouter("/updateUserInfo", &api.UserInfoController{}, "post:UpdateUserInfo"),
		),
		beego.NSNamespace("/images",
			beego.NSRouter("/uploadFiles", &api.ImageController{}, "post:UploadFiles")),
		beego.NSNamespace("/merchant",
			beego.NSRouter("/merchantList", &api.MerchantController{}, "post:MerchantList"),
			beego.NSRouter("/merchantDetail", &api.MerchantController{}, "post:MerchantDetail")),
	)

	beego.AddNamespace(apiPath)
}
