package api

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
	"my-ganji-app/models"
	"my-ganji-app/types"
	"os"
	"path"
	"strings"
	"time"
)

type ImageController struct {
	beego.Controller
}

type Sizer interface {
	Size() int64
}

// UploadFiles 实现上传图片功能
// 实现本地预览的流程：
// 1. 保存到本地目录 upload_image
// 2. 返回可返回的 URL
// 3. Beego SetStaticPath 把 /static 映射到该目录
// 4. 浏览器打开 URL 即可预览
func (c *ImageController) UploadFiles() {
	f, h, err := c.GetFile("file")
	if err != nil {
		c.Data["json"] = RetResource(false, types.GetImagesFileFail, nil, "获取文件失败")
		c.ServeJSON()
		return
	}
	defer f.Close()

	// 文件格式校验
	ext := strings.ToLower(path.Ext(h.Filename))
	allowExtMap := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true}
	if !allowExtMap[ext] {
		c.Data["json"] = RetResource(false, types.FileFormatError, nil, "上传的文件格式不符合要求")
		c.ServeJSON()
		return
	}

	// 文件大小相关
	const maxFileBytes int64 = 1 << 24 //  16MB
	fileSizer, ok := f.(Sizer)
	if !ok {
		c.Data["json"] = RetResource(false, types.GetImagesFileFail, nil, "无法获取文件大小")
		c.ServeJSON()
		return
	}

	if fileSizer.Size() > maxFileBytes {
		c.Data["json"] = RetResource(false, types.FileIsBig, nil, "文件太大了，最大支持16MB")
		c.ServeJSON()
		return
	}

	frontImagePath, err := beego.AppConfig.String("front_image_path")
	if err != nil {
		c.Data["json"] = RetResource(false, types.InvalidConfig, nil, "没有配置 front_image_path 环境变量")
	}
	imgDir, err := beego.AppConfig.String("upload_image")
	if err != nil {
		c.Data["json"] = RetResource(false, types.InvalidConfig, nil, "没有配置 upload_image 环境变量")
	}
	timeStr := time.Now().Format("2006/01/02/")
	uploadDir := imgDir + timeStr
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.Data["json"] = RetResource(false, types.CreateFilePathError, nil, "文件路径创建失败")
		c.ServeJSON()
		return
	}

	// 用 UUID 重命名，避免同名覆盖
	fileName := uuid.New().String() + ext
	fpath := uploadDir + fileName
	if err = c.SaveToFile("file", fpath); err != nil {
		c.Data["json"] = RetResource(false, types.SaveFileFail, err.Error(), "保存文件失败")
		c.ServeJSON()
		return
	}

	imgUrl := frontImagePath + timeStr + fileName
	imgFile := models.ImageFile{Url: imgUrl}
	err, id := imgFile.Insert()
	if err != nil {
		_ = os.Remove(fpath)
		c.Data["json"] = RetResource(false, types.FileAlreadyUpload, nil, "图片保存失败，请重试")
		c.ServeJSON()
		return
	}
	data := map[string]interface{}{
		"image_id": id,
		"img_url":  imgUrl,
	}
	c.Data["json"] = RetResource(true, types.ReturnSuccess, data, "上传文件成功")
	c.ServeJSON()
}
