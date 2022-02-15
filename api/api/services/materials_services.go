package services

import (
	"cramee/myerror"
	"cramee/token"
	"cramee/util"
	"fmt"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type MaterialsService interface {
	UploadMaterials(file multipart.FileHeader, status string, teacherId string, studentId string, schoolHours string) error
	GetUrlOfMarterials(status string, teacherId string, studentId string, schoolHours string) (string, error)
}

type materialsServiceImpl struct {
	config     util.Config
	tokenMaker token.Maker
}

func NewMaterialsService(config util.Config, tokenMaker token.Maker) MaterialsService {
	res := &materialsServiceImpl{}
	res.config = config
	res.tokenMaker = tokenMaker
	return res
}

func (m *materialsServiceImpl) UploadMaterials(file multipart.FileHeader, status string, teacherId string, studentId string, schoolHours string) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(m.config.AwsS3AccessKey, m.config.AwsS3SecretAccessKey, ""),
		Region:      aws.String(m.config.AwsS3Region),
	}))
	uploader := s3manager.NewUploader(sess)
	key := status + "/" + teacherId + "/" + studentId + "/" + schoolHours
	data, err := os.Open(file.Filename)
	if err != nil {
		return myerror.NewPublic(myerror.ErrUpload, err)
	}
	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(m.config.AwsS3Bucket),
		Key:    aws.String(key),
		Body:   data,
	})
	if err != nil {
		fmt.Println(res)
		if err, ok := err.(awserr.Error); ok && err.Code() == request.CanceledErrorCode {
			return myerror.NewPublic(myerror.ErrTimeOut, err)
		} else {
			return myerror.NewPublic(myerror.ErrUpload, err)
		}
	}
	defer data.Close()
	return nil
}

func (m *materialsServiceImpl) GetUrlOfMarterials(status string, teacherId string, studentId string, schoolHours string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(m.config.AwsS3AccessKey, m.config.AwsS3SecretAccessKey, ""),
		Region:      aws.String(m.config.AwsS3Region)},
	)
	if err != nil {
		return "", myerror.NewPublic(myerror.ErrBindData, err)
	}
	svc := s3.New(sess)
	key := status + "/" + teacherId + "/" + studentId + "/" + schoolHours
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(m.config.AwsS3Bucket),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(168 * time.Minute)
	if err != nil {
		return "", myerror.NewPublic(myerror.ErrBindData, err)
	}
	return urlStr, nil
}
