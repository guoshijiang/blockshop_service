package api

import (
	"blockshop/models"
	"blockshop/types"
	"github.com/astaxie/beego"
	"os"
	"path"
	"time"
)

type ImageController struct {
	beego.Controller
}

type Sizer interface {
	Size() int64
}

// UploadFiles @Title UploadFiles
// @Description 上传图片 UploadFiles
// @Success 200 status bool, data interface{}, msg string
// @router /upload_file [post]
func (this *ImageController) UploadFiles() {
	f, h, err := this.GetFile("file")
	if err != nil {
		this.Data["json"] = RetResource(false, types.GetImagesFileFail, nil, "获取文件失败")
		this.ServeJSON()
		return
	}
	defer f.Close()
	ext := path.Ext(h.Filename)
	var AllowExtMap map[string]bool = map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true}
	if _, ok := AllowExtMap[ext]; !ok {
		this.Data["json"] = RetResource(false, types.FileFormatError, nil, "上传的文件格式不符合要求")
		this.ServeJSON()
		return
	}
	var Filebytes = 1 << 24
	if fileSizer, ok := f.(Sizer); ok {
		fileSize := fileSizer.Size()
		if fileSize > int64(Filebytes) {
			this.Data["json"] = RetResource(false, types.FileIsBig, nil, "文件太大了")
			this.ServeJSON()
		} else {
			front_image_path := beego.AppConfig.String("front_image_path")
			img_dir := beego.AppConfig.String("upload_image")
			time_str := time.Now().Format("2006/01/02/")
			uploadDir := img_dir + time_str
			err = os.MkdirAll(uploadDir, os.ModePerm)
			if err != nil {
				this.Data["json"] = RetResource(false, types.CreateFilePathError, nil, "文件路径创建失败")
				this.ServeJSON()
				return
			}
			fpath := uploadDir + h.Filename
			err = this.SaveToFile("file", fpath)
			if err != nil {
				this.Data["json"] = RetResource(false, types.FileIsBig, err.Error(), "保存文件失败")
				this.ServeJSON()
				return
			}
			img_file := models.ImageFile{
				Url: front_image_path + time_str + h.Filename,
			}
			err, id := img_file.Insert()
			if err != nil {
				this.Data["json"] = RetResource(true, types.FileAlreadUpload, nil, "该图片已经上传过了")
				this.ServeJSON()
				return
			}
			data := map[string]interface{}{
				"imgage_id": id,
				"img_url": front_image_path + time_str + h.Filename,
			}
			this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "上传文件成功")
			this.ServeJSON()
			return
		}
	} else {
		this.Data["json"] = RetResource(false, types.FileIsBig, nil, "上传文件太大")
		this.ServeJSON()
		return
	}
}
