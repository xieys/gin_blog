package model

import (
	"context"
	"gin_blog/utils"
	"gin_blog/utils/errmsg"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

var (
	accessKey = utils.AccessKey
	secretKey = utils.SecretKey
	bucket    = utils.Bucket
	imgUrl    = utils.QiniuSever
)

func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong, // 空间对应的机房
		UseHTTPS:      false,                // 是否使用https域名
		UseCdnDomains: false,                // 上传是否使用CDN上传加速
	}
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := imgUrl + ret.Key
	return url, errmsg.SUCCESS
}
