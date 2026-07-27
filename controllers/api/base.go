package api

import (
	beego "github.com/beego/beego/v2/server/web"
	"my-ganji-app/models"
	"my-ganji-app/types"
	"strings"
)

const HttpAuthKey = "Authorization"

type RetJson struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

func RetResource(status bool, code int, data interface{}, msg string) (apijson *RetJson) {
	apijson = &RetJson{Status: status, Code: code, Data: data, Msg: msg}
	return
}

type BaseController struct {
	beego.Controller
}

// CurrentUser 提取当前登陆用户
func (c *BaseController) CurrentUser() (*models.User, bool) {
	bearerToken := c.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		c.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		c.ServeJSON()
		return nil, false
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user, err := models.GetUserByToken(token)
	if err != nil {
		c.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		c.ServeJSON()
		return nil, false
	}
	return user, true
}
