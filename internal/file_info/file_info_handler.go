package file_info

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Summary Upload a file image
// @Description Upload a file image and store information
// @Accept mpfd
// @Produce json
// @Tags Files
// @Param file formData file true "File to upload"
// @Param fileName formData string false "File name"
// @Param description formData string false "File description"
// @Param infoType formData string true "Information type" Enums(demo/characters, demo/animes)
// @Param updateBy formData string true "Updated by"
// @Success 200 {object} FileInfo "Successful response"
// @Router /files/uploadFileImage [post]
func UploadFileImage(c *gin.Context) {
	fmt.Println("UploadFileImage start....")
	fileInfoService := &FileInfoService{}
	fileInfoService.UploadFileImage(c)
	fmt.Println("UploadFileImage success....")
}
