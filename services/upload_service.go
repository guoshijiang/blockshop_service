package services

import (
  "blockshop/models"
  "errors"
  "fmt"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/context"
  "github.com/astaxie/beego/orm"
  "github.com/google/uuid"
  "io"
  "mime/multipart"
  "net/http"
  "os"
  "path"
  "path/filepath"
  "strconv"
  "strings"
)

type UploadService struct {
}

//上传单个文件
func (*UploadService) Upload(ctx *context.Context, name string) (string,error) {
  file, h, err := ctx.Request.FormFile(name)
  if err != nil {
    fmt.Println("err---",err)
    return "",err
  }
  defer file.Close()

  //自定义文件验证
  err = validateForUpload(h)
  if err != nil {
    return "",err
  }

  //数据表写入
  saveName := uuid.New().String()
  //后缀带. (.png)
  fileExt := path.Ext(h.Filename)
  savePath := beego.AppConfig.String("images::path") + saveName + fileExt
  saveRealDir := filepath.ToSlash(beego.AppPath + "/" + beego.AppConfig.String("images::path"))

  _, err = os.Stat(saveRealDir)
  if err != nil {
    err = os.MkdirAll(saveRealDir, os.ModePerm)
  }

  saveUrl := "/" + beego.AppConfig.String("images::url") + saveName + fileExt
  f, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
  if err != nil {
    return "", err
  }
  defer f.Close()
  io.Copy(f, file)
  return saveUrl,nil
}

//上传多个文件
func (*UploadService) UploadMulti(ctx *context.Context, name string, GoodsId int64) ([]*models.GoodsImage, error) {
  var (
    result []*models.GoodsImage
    data []*models.GoodsImage
  )
  files, ok := ctx.Request.MultipartForm.File[name]
  if !ok {
    return nil, http.ErrMissingFile
  }
  //清除
  orm.NewOrm().QueryTable(new(models.GoodsImage)).Filter("goods_id__eq",GoodsId).All(&data)
  for _,v := range data {
    err := os.Remove(v.Image[1:])
    fmt.Println("remove_err",err)
  }
  orm.NewOrm().QueryTable(new(models.GoodsImage)).Filter("goods_id__eq",GoodsId).Delete()

  for i, _ := range files {
    h := files[i]
    file, err := files[i].Open()
    defer file.Close()
    if err != nil {
      return nil, err
    }
    err = validateForUpload(h)
    if err != nil {
      return nil, err
    }

    //数据表写入
    saveName := uuid.New().String()
    //后缀带. (.png)
    fileExt := path.Ext(h.Filename)
    savePath := beego.AppConfig.String("images::path") + saveName + fileExt
    saveRealDir := filepath.ToSlash(beego.AppPath + "/" + beego.AppConfig.String("images::path"))

    _, err = os.Stat(saveRealDir)
    if err != nil {
      err = os.MkdirAll(saveRealDir, os.ModePerm)
    }

    saveUrl := "/" + beego.AppConfig.String("images::url") + saveName + fileExt

    f, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
      return nil, err
    }
    defer f.Close()
    io.Copy(f, file)

    goodsImagesInfo := models.GoodsImage{
      GoodsId:  GoodsId,
      Image:  saveUrl,
      IsShow:int8(1),
    }

    insertId, err := orm.NewOrm().Insert(&goodsImagesInfo)
    if err == nil {
      goodsImagesInfo.Id = int64(insertId)
      result = append(result, &goodsImagesInfo)
    } else {
      return nil, err
    }
  }
  //返回上传文件的对象信息
  if result != nil {
    return result, nil
  } else {
    return nil, errors.New("无法获取文件")
  }

}

//images自定义验证
func validateForUpload(h *multipart.FileHeader) error {
  validateSize, _ := strconv.Atoi(beego.AppConfig.String("images::validate_size"))
  validateExt := beego.AppConfig.String("images::validate_ext")
  if int(h.Size) > validateSize {
    return errors.New("文件超过限制大小")
  }

  if !strings.Contains(validateExt, strings.ToLower(strings.TrimLeft(path.Ext(h.Filename), "."))) {
    return errors.New("不支持的文件格式")
  }
  return nil
}