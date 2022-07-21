package routes

import (
	"backend/configs"
	"backend/tools"
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"io/ioutil"
	"strings"
)

func DownloadFile(c echo.Context) error {

	url := c.Request().RequestURI
	url_list := strings.Split(url, "/")
	filename := url_list[len(url_list)-1]

	client, ctx := configs.GetMongoClient()
	defer client.Disconnect(ctx)

	db := client.Database("test")
	bucket, _ := gridfs.NewBucket(
		db,
	)
	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(filename, &buf)
	tools.CheckError(err)

	fmt.Printf("File size to download: %v \n", dStream)
	ioutil.WriteFile(filename, buf.Bytes(), 0600)

	return c.File(filename)
}
