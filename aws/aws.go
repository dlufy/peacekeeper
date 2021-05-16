package aws

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/dlufy/peacekeeper/admin"
	"github.com/rs/xid"
)

var (
	sess       *session.Session
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
)

func Initialize() error {
	mainCfg := admin.GetConfig()
	cred := credentials.NewStaticCredentials(mainCfg.AWS.AcessKey, mainCfg.AWS.SecretAcessKey, "")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(mainCfg.AWS.Region),
		Credentials: cred,
	})
	if err != nil {
		log.Printf("error while setting session for s3 err %+v", err)
		return err
	}
	uploader = s3manager.NewUploader(sess)
	downloader = s3manager.NewDownloader(sess)
	return err
}
func UploadFileToS3(file io.Reader, bucketName, Key, contentType, ACL string) (string, error) {
	if uploader == nil {
		Initialize()
		if uploader == nil {
			log.Printf("not able to initialize the uploader\n")
			return "", errors.New("uploader is not initalized")
		}
	}
	if Key == "" {
		Key = "temp_" + xid.New().String()
	}
	if contentType == "" {
		contentType = "image/png"
	}
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(Key),
		Body:        file,
		ContentType: &contentType,
		ACL:         &ACL,
	})

	if err != nil {
		log.Printf("Error while uploading file to s3 %+v\n", err)
		return "", err
	}
	log.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return result.Location, nil
}

func DownloadFileFromS3(fileURL, fileName, dir, bucketName string) (string, error) {
	key, err := GetKeyFromS3URL(fileURL, dir)
	// Create a file to write the S3 Object contents to.
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("failed to create file %q, %v\n", fileName, err)
		return "", err
	}

	// Write the contents of S3 Object to the file
	n, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	log.Printf("file downloaded successfully bytes read %d\n", n)
	return fileName, err
}

func GetKeyFromS3URL(s3URL, dir string) (string, error) {
	cfg := admin.GetConfig()
	if !strings.Contains(s3URL, cfg.AWS.BucketName) {
		return "", errors.New("invalid s3 url")
	}
	arr := strings.Split(s3URL, dir)
	if len(arr) != 2 {
		return "", errors.New("invalid s3url")
	}
	return arr[1], nil
}
