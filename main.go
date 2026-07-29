package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "my-ganji-app/routers"
)

func main() {
	// 把本地上传目录映射到 /static，对应 front_image_path
	uploadDir, err := beego.AppConfig.String("upload_image")
	if err != nil || uploadDir == "" {
		uploadDir = "static"
	}
	beego.SetStaticPath("/static", uploadDir)

	beego.Run("127.0.0.1:8080")
}
