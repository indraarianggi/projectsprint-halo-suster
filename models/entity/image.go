package entity

import "mime/multipart"

type UploadFileRequest struct {
	File        *multipart.FileHeader `json:"file" form:"file" validate:"required"`
	ContentType string                `json:"content_type" form:"content_type"`
}

type UploadS3Response struct {
	Size     int
	Mimetype string
	Name     string
	Location string
}

type UploadImageResponse struct {
	Message string          `json:"message"`
	Data    UploadImageData `json:"data"`
}

type UploadImageData struct {
	ImageURL string `json:"imageUrl"`
}
