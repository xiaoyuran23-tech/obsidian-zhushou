package store

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	*s3.S3
	bucket string
}

func InitS3(endpoint, accessKey, secretKey, region string) (*S3Client, error) {
	if endpoint == "" {
		return nil, fmt.Errorf("S3 endpoint is required")
	}

	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(region),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	return &S3Client{
		S3:     s3.New(sess),
		bucket: "obsidian-zhushou",
	}, nil
}

func (c *S3Client) PutObject(key string, data []byte) error {
	_, err := c.S3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
		Body:   aws.ReadSeekCloser(data),
	})
	return err
}

func (c *S3Client) GetObject(key string) ([]byte, error) {
	result, err := c.S3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	buf := make([]byte, *result.ContentLength)
	n, err := result.Body.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func (c *S3Client) DeleteObject(key string) error {
	_, err := c.S3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
	})
	return err
}
