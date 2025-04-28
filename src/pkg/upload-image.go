package pkg

import (
	"BE_Ecommerce/src/config"
	"context"
	"log"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func newCloudInstance() (*cloudinary.Cloudinary, error) {
	config := config.LoadEnv()
	cloud, err := cloudinary.NewFromParams(config.CloudName, config.CloudApiKey, config.CloudApiSecret)
	if err != nil {
		log.Println("Failed to create cloudinary instance", err.Error())
	}
	return cloud, nil
}

func UploadSingleImage(fileHeader *multipart.FileHeader, folder string) (*string, error) {
	cloud, err := newCloudInstance()
	if err != nil {
		return nil, err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}

	defer file.Close()

	uploadResult, err := cloud.Upload.Upload(context.Background(), file, uploader.UploadParams{Folder: folder})
	if err != nil {
		return nil, err
	}
	return &uploadResult.URL, nil
}

func UploadMultipleImages(fileHeaders []*multipart.FileHeader, folder string) ([]string, error) {
	cloud, err := newCloudInstance()
	if err != nil {
		return nil, err
	}

	var urls []string
	for _, fileHeader := range fileHeaders {

		file, err := fileHeader.Open()

		if err != nil {
			return nil, err
		}
		defer file.Close()

		uploadResult, err := cloud.Upload.Upload(context.Background(), file, uploader.UploadParams{Folder: folder})
		if err != nil {
			return nil, err
		}
		urls = append(urls, uploadResult.URL)
	}
	return urls, nil
}
