package files

//
//import (
//	"errors"
//	"github.com/aliyun/aliyun-oss-go-sdk/oss"
//	"mime/multipart"
//	"path/filepath"
//	"time"
//)
//
//type AliOSS struct{
//
//}
//
//func NewBucket() (*oss.Bucket, error) {
//
//	endpoint := variable.ConfigYml.GetString("AliYun.Oss.Endpoint")
//	accessKeyID := variable.ConfigYml.GetString("AliYun.Oss.AccessKeyID")
//	accessKeySecret := variable.ConfigYml.GetString("AliYun.Oss.AccessKeySecret")
//	bucketConfig := variable.ConfigYml.GetString("AliYun.Oss.Bucket")
//	// 创建OSSClient实例
//	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
//	if err != nil {
//		return nil, err
//	}
//	bucket, err := client.Bucket(bucketConfig)
//	if err != nil {
//		return nil, err
//	}
//
//	return bucket, nil
//}
//
//
//func (*AliOSS) UploadFile(file *multipart.FileHeader) (string, string, error) {
//	bucket, err := NewBucket()
//	if err != nil {
//		return "", "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
//	}
//
//	// 读取本地文件。
//	f, openError := file.Open()
//	if openError != nil {
//		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
//	}
//	defer func(f multipart.File) {
//		err := f.Close()
//		if err != nil {
//
//		}
//	}(f) // 创建文件 defer 关闭
//	// 上传阿里云路径 文件名格式 自己可以改 建议保证唯一性
//	yunFileTmpPath := filepath.Join("uploads", time.Now().Format("2006-01-02")) + "/" + file.Filename
//
//	// 上传文件流。
//	err = bucket.PutObject(yunFileTmpPath, f)
//	if err != nil {
//		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
//	}
//	return variable.ConfigYml.GetString("AliYun.Oss.CdnDomain") + "/" + yunFileTmpPath, yunFileTmpPath, nil
//
//}
//
//func (*AliOSS) DeleteFile(key string) error {
//	bucket, err := NewBucket()
//	if err != nil {
//		//global.GVA_LOG.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
//		return errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
//	}
//	err = bucket.DeleteObject(key)
//	if err != nil {
//		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
//	}
//	return nil
//}
//
