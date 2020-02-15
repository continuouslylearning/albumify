package db

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

type S3Handler struct {
	Session *session.Session
	Bucket  string
}

func InitializeS3Handler(r *gin.Engine) *S3Handler {
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("S3_BUCKET")

	session, _ := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	h := &S3Handler{
		Session: session,
		Bucket:  bucket,
	}
	r.Use(AddS3Handler(h))

	return h
}

func AddS3Handler(h *S3Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("s3", h)
		c.Next()
	}
}

func (h S3Handler) DeleteImage(key string, username string) error {
	_, e := s3.New(h.Session).DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(h.Bucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", username, key)),
	})

	return e
}

func (h S3Handler) GetAlbumKeys(username string) ([]string, error) {
	res, e := s3.New(h.Session).ListObjects(&s3.ListObjectsInput{
		Bucket:    aws.String(h.Bucket),
		Delimiter: aws.String("/"),
		MaxKeys:   aws.Int64(18),
		Prefix:    aws.String(fmt.Sprintf("%s/", username)),
	})
	if e != nil {
		return nil, e
	}

	var keys []string
	for _, object := range res.Contents {
		keys = append(keys, *object.Key)
	}

	return keys, nil
}

func (h S3Handler) GetURLs(username string, keys []string, imageURLs []string) ([]string, error) {
	for i, key := range keys {
		if len(imageURLs[i]) == 0 {
			req, _ := s3.New(h.Session).GetObjectRequest(&s3.GetObjectInput{
				Bucket: aws.String(h.Bucket),
				Key:    aws.String(key),
			})

			signedURL, e := req.Presign(60 * time.Minute)
			if e != nil {
				return nil, e
			}

			imageURLs[i] = signedURL
		}
	}
	return imageURLs, nil
}

func (h S3Handler) UploadImage(key string, body []byte, username string) error {
	_, e := s3.New(h.Session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(h.Bucket),
		Key:                  aws.String(fmt.Sprintf("%s/%s", username, key)),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(body),
		ContentLength:        aws.Int64(int64(len(body))),
		ContentType:          aws.String(http.DetectContentType(body)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return e
}

func (h S3Handler) UploadImages(keys []string, images [][]byte, username string) error {
	var objects []s3manager.BatchUploadObject

	for i, key := range keys {
		body := images[i]
		objects = append(objects, s3manager.BatchUploadObject{
			Object: &s3manager.UploadInput{
				Bucket:             aws.String(h.Bucket),
				Key:                aws.String(fmt.Sprintf("%s/%s", username, key)),
				ACL:                aws.String("private"),
				Body:               bytes.NewReader(body),
				ContentDisposition: aws.String("attachment"),
				ContentType:        aws.String(http.DetectContentType(body)),
			},
		})
	}

	iter := &s3manager.UploadObjectsIterator{
		Objects: objects,
	}

	return s3manager.NewUploader(h.Session).UploadWithIterator(aws.BackgroundContext(), iter)
}
