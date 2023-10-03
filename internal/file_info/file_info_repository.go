package file_info

import (
	"fmt"
	"go-demo/config"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileInfoRepository struct {
}

func (*FileInfoRepository) UploadFileImage(c *gin.Context) (string, string, string, error) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println("Error parsing the multipart form:", err)
		return "", "", "", err
	}

	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println("Error retrieving the file:", err)
		return "", "", "", err
	}
	defer file.Close()

	filename := handler.Filename
	extension := filepath.Ext(filename)

	// บันทึกไฟล์ลงในโฟลเดอร์ uploads
	uploadPath := "uploads"
	err = os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating upload directory:", err)
		return "", "", "", err
	}

	filePath := filepath.Join(uploadPath, filename)
	err = c.SaveUploadedFile(handler, filePath)
	if err != nil {
		fmt.Println("Error saving the file:", err)
		return "", "", "", err
	}

	fileURL := fmt.Sprintf("http://localhost:80/uploads/%s", filename)

	fmt.Printf("File uploaded successfully. URL: %s\n", fileURL)
	return filename, fileURL, extension, nil
}

func (r *FileInfoRepository) DBCreateFile(request FileInfoRequest, fileName string, extension string, url string) (*FileInfo, error) {
	characterID := uuid.New()
	fileInfo := FileInfo{
		ID:          characterID,
		FileName:    fileName,
		Extension:   extension,
		Url:         url,
		Description: request.Description,
		InfoType:    request.InfoType,
		Activate:    true,
		CreateBy:    request.UpdateBy,
		UpdateBy:    request.UpdateBy,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
	config.DB.Create(&fileInfo)
	return &fileInfo, nil
}
