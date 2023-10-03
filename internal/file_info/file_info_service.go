package file_info

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileInfoService struct{}

func (s *FileInfoService) UploadFileImage(c *gin.Context) {
	var inputFileInfoRequest FileInfoRequest
	if err := c.ShouldBind(&inputFileInfoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("UploadFileImage service...")

	fileInfoRepository := &FileInfoRepository{}
	filename, fileURL, extension, err := fileInfoRepository.UploadFileImage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fileInfo, err := fileInfoRepository.DBCreateFile(inputFileInfoRequest, filename, extension, fileURL)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("UploadFileImage Service success...")

	c.JSON(http.StatusCreated, fileInfo)
}
