package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "my-ganji-app/routers"
)

func main() {
	beego.Run("127.0.0.1:8080")
}
