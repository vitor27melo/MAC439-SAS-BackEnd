package routes

import (
	"backend/configs"
	"backend/tools"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func DownloadFile(c echo.Context) error {
	client, ctx := configs.GetMongoClient()
	defer client.Disconnect(ctx)

	db := client.Database("test")
	bucket, _ := gridfs.NewBucket(
		db,
	)
	var buf bytes.Buffer
	filename := "20220715051349download.jpeg"
	dStream, err := bucket.DownloadToStreamByName(filename, &buf)
	tools.CheckError(err)

	fmt.Printf("File size to download: %v \n", dStream)
	ioutil.WriteFile(filename, buf.Bytes(), 0600)

	return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("File %s downloaded successfully ", filename)})
}
