package services

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/xo-yosi/Talent-SMPS/internal/app/models"
	"github.com/xo-yosi/Talent-SMPS/internal/app/repository"
	"github.com/xo-yosi/Talent-SMPS/internal/config"
	"github.com/xo-yosi/Talent-SMPS/internal/utils"
)

type StudentService struct {
	srepo    repository.StudentRepository
	S3Client *s3.Client
}

func NewStudentService(srepo repository.StudentRepository, s3Client *s3.Client) *StudentService {
	return &StudentService{srepo: srepo, S3Client: s3Client}
}

func (s *StudentService) RegisterStudent(studentData models.StudentRegisterRequest, file *multipart.FileHeader) (int, error) {
	lastID, err := s.srepo.GetLastCoordinatorID()
	if err != nil {
		return 0, errors.New("failed to get last Student ID")
	}

	newID := utils.GenerateNextCoordinatorID(lastID)
	var profilePicURL string
	bucket := "talent-smps-images"
	if file != nil {
		src, err := file.Open()
		if err != nil {
			return 0, fmt.Errorf("failed to open profile picture: %w", err)
		}
		defer src.Close()

		key := fmt.Sprintf("profile-pics/%s_%s", studentData.Name, filepath.Base(file.Filename))

		_, err = s.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
			Body:   src,
			ACL:    "public-read",
		})
		if err != nil {
			return 0, fmt.Errorf("failed to upload profile picture: %w", err)
		}

		profilePicURL = fmt.Sprintf("%s/%s/%s", config.AppConfig.S3Endpoint, bucket, key)
	}

	student := models.Student{
		StudentID:   newID,
		Name:        studentData.Name,
		Age:         studentData.Age,
		PhoneNumber: studentData.PhoneNumber,
		Gender:      studentData.Gender,
		ProfilePic:  profilePicURL,
		CreatedAt:   time.Now(),
	}

	if err := s.srepo.CreateStudent(&student); err != nil {
		return 0, errors.New("failed to create student")
	}

	return newID, nil
}
