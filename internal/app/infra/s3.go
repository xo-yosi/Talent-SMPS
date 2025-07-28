package infra

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/xo-yosi/Talent-SMPS/internal/config"
)

func NewClient() *s3.Client {
	return s3.NewFromConfig(aws.Config{
		Region: config.AppConfig.S3Region,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
			config.AppConfig.S3Endpoint, config.AppConfig.S3SecretKey, "",
		)),
		EndpointResolverWithOptions: aws.EndpointResolverWithOptionsFunc(
			func(service, region string, _ ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           config.AppConfig.S3Endpoint,
					SigningRegion: config.AppConfig.S3Region,
				}, nil
			}),
	})
}
