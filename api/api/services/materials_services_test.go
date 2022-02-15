package services

import (
	"cramee/util"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func TestUploadMarterials(t *testing.T) {
	testcases := []struct {
		name      string
		teacherId string
		studentId string
		file      string
	}{
		{
			name:      "test-case-no.00",
			teacherId: "de94e2739851a71eb99697f9d65b10c77886e9c222d03a0f57b2073d4ad8c7dd",
			studentId: "8567ec64e98fd5ddca01929879da1114d515a87c5fa4308ba9a94460eaf54612",
			file:      "./test-case/test-case-00.pdf",
		},
		{
			name:      "test-case-no.01",
			teacherId: "3df3b6520cd8592fe6b121284d0f1984655418676e75c741fe23b5c8ad0b3743",
			studentId: "176c86b942bad1630a816297ccc30022e1f5b0042ce7b5276eabd00d031c4612",
			file:      "./test-case/test-case-01.pdf",
		},
		{
			name:      "test-case-no.02",
			teacherId: "46236b23e915b70d8afe4aa3bebd96c226e141756e4a47fc0f171e7a69411a19",
			studentId: "55d49be1b443298dfa8a612228ed151b2d6ecf9785a80d3c17891f60f7af7d67",
			file:      "./test-case/test-case-02.pdf",
		},
		{
			name:      "test-case-no.03",
			teacherId: "c7201eaf1c51146972a4055f17376273efe6637bf4a3a1a579df933636b64d30",
			studentId: "d6dd90a5939fd82c974be05abb80ff61e2ad70f1f39162794392b387ff1b7c73",
			file:      "./test-case/test-case-03.pdf",
		},
		{
			name:      "test-case-no.04",
			teacherId: "73f205333aeb544ab00d61587346946f066a7abee80aec71516021967e7a9535",
			studentId: "552566866442d8a41fe7e0cdef7cef05d0f1214f45d8ebafccbb1f5cb090b9c9",
			file:      "./test-case/test-case-04.pdf",
		},
		{
			name:      "test-case-no.05",
			teacherId: "df6d7e2e786e9497fc5b2fea210dea22cafeebd001d9d8612b51adbbaa1694ff",
			studentId: "4dc4635073051921e4ba83432cd4734c58c792a0076824751d80e02fc0755c9f",
			file:      "./test-case/test-case-05.pdf",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			file, err := os.Open(testcase.file)
			if err != nil {
				t.Errorf("could not open file")
			}
			config, err := util.LoadConfig("../..")
			if err != nil {
				t.Errorf("config error")
			}
			sess := session.Must(session.NewSession(&aws.Config{
				Credentials: credentials.NewStaticCredentials(config.AwsS3AccessKey, config.AwsS3SecretAccessKey, ""),
				Region:      aws.String(config.AwsS3Region),
			}))
			uploader := s3manager.NewUploader(sess)
			key := "teacher/" + testcase.teacherId + "/" + testcase.studentId + "/" + time.Now().Local().Format("2006-01-02--00-00-00")
			_, err = uploader.Upload(&s3manager.UploadInput{
				Bucket: aws.String(config.AwsS3Bucket),
				Key:    aws.String(key),
				Body:   file,
			})
			if err != nil {
				if err, ok := err.(awserr.Error); ok && err.Code() == request.CanceledErrorCode {
					t.Errorf("time out")
				} else {
					t.Errorf("could not upload file")
				}
			}
		})
	}
}
