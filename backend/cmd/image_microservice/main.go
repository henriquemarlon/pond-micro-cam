package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/henriquemarlon/pond-micro-cam/backend/configs"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
	"net/http"
)

func main() {

	/////////////////////// Configs /////////////////////////

	db := configs.SetupPostgres()
	defer db.Close()

	/////////////////////////////////////////////////////////

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	api := router.Group("/image")
	api.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Get form err: %s", err.Error()))
			return
		}
		src, err := file.Open()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("File open err: %s", err.Error()))
			return
		}
		defer src.Close()
		img, _, err := image.Decode(src)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Image decode err: %s", err.Error()))
			return
		}
		grayImg := image.NewGray(img.Bounds())
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				originalColor := img.At(x, y)
				grayColor := color.GrayModel.Convert(originalColor)
				grayImg.Set(x, y, grayColor)
			}
		}
		buf := new(bytes.Buffer)
		if err := jpeg.Encode(buf, grayImg, nil); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("JPEG encode err: %s", err.Error()))
			return
		}
		base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
		c.String(http.StatusOK, base64Str)
	})

	router.Run(":8081")
}
