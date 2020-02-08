package database

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
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

func (h S3Handler) UploadImage(key string, body []byte) error {
	_, e := s3.New(h.Session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(h.Bucket),
		Key:                  aws.String("twice/" + key),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(body),
		ContentLength:        aws.Int64(int64(len(body))),
		ContentType:          aws.String(http.DetectContentType(body)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return e
}

func (h S3Handler) GetImages() ([]string, error) {
	res, e := s3.New(h.Session).ListObjects(&s3.ListObjectsInput{
		Bucket:    aws.String(h.Bucket),
		Delimiter: aws.String("/"),
		MaxKeys:   aws.Int64(18),
		Prefix:    aws.String("twice/"),
	})

	if e != nil {
		return nil, e
	}

	var keys []string
	for _, object := range res.Contents {
		keys = append(keys, *object.Key)
	}

	images := make([]string, 0, len(keys))
	for _, key := range keys {
		req, _ := s3.New(h.Session).GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(h.Bucket),
			Key:    aws.String(key),
		})
		signedURL, _ := req.Presign(60 * time.Minute)
		images = append(images, signedURL)
	}

	return images, nil
}
