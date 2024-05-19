package usecase

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/backend-magang/halo-suster/models/entity"
	"github.com/backend-magang/halo-suster/utils/constant"
	"github.com/backend-magang/halo-suster/utils/helper"
)

func (u *usecase) UploadImage(ctx context.Context, fileHeader *multipart.FileHeader) helper.StandardResponse {
	if err := validateFile(fileHeader); err != nil {
		return helper.StandardResponse{Code: http.StatusBadRequest, Message: constant.FAILED, Error: err}
	}

	file, err := fileHeader.Open()
	if err != nil {
		u.logger.Errorf("[Usecase][Image][UploadImage] failed to open file, err: %s", err.Error())
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}
	defer file.Close()

	fileName := generateUniqueFileName()

	obj := &s3.PutObjectInput{
		Bucket: aws.String(u.config.S3BucketName),
		Key:    aws.String(fileName),
		ACL:    types.ObjectCannedACLPublicRead,
		Body:   file,
	}

	_, err = u.s3.PutObject(ctx, obj)
	if err != nil {
		u.logger.Errorf("[Usecase][Image][UploadImage] failed to put obecjt s3, err: %s", err.Error())
		return helper.StandardResponse{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
	}

	imageURL := fmt.Sprintf("https://%s.s3.ap-southeast-1.amazonaws.com/%s", u.config.S3BucketName, fileName)
	uploadData := entity.UploadImageData{ImageURL: imageURL}

	return helper.StandardResponse{Code: http.StatusOK, Message: constant.SUCCESS_UPLOAD_FILE, Data: uploadData}
}
