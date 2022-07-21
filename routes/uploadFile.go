package routes

import (
	"backend/configs"
	"backend/tools"
	"database/sql"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"io/ioutil"
	"net/http"
	"time"
)

type Attachment struct {
	AttachmentType string `json:"attachmentType"`
	Date           string `json:"date"`
	Observation    string `json:"obs"`
}

var stmtQuery = `
	SELECT
		cpf
	FROM
		usuario
	WHERE
		id_usuario = $1;
`

func UploadFile(c echo.Context) error {
	attachment := new(Attachment)
	attachment.AttachmentType = c.FormValue("attachmentType")
	attachment.Date = c.FormValue("date")
	attachment.Observation = c.FormValue("obs")
	file, err := c.FormFile("file")

	tools.CheckError(err)
	filename := file.Filename
	fileContent, err := file.Open()
	t := time.Now().Format("20060102150405")
	attachmentName := t + "-filebegin-" + filename
	tools.CheckError(err)

	fileData, err := ioutil.ReadAll(fileContent)
	tools.CheckError(err)
	fileType := mimetype.Detect(fileData).String()

	client, ctx := configs.GetMongoClient()
	defer client.Disconnect(ctx)

	bucket, err := gridfs.NewBucket(client.Database("test"))
	tools.CheckError(err)

	uploadStream, err := bucket.OpenUploadStream(attachmentName)
	tools.CheckError(err)
	defer uploadStream.Close()

	fileSize, _ := uploadStream.Write(fileData)
	if fileSize == 0 {
		return c.JSON(http.StatusUnsupportedMediaType, map[string]string{"message": "Invalid file size error"})
	}

	db, err := sql.Open(configs.GetDBType(), configs.GetPostgresConnString())
	tools.CheckError(err)

	defer db.Close()

	userId := c.Get("userId")
	var cpf string
	err = db.QueryRow(stmtQuery, userId).Scan(&cpf)
	tools.CheckError(err)

	coll := client.Database("test").Collection("usuario")
	doc := prepareDoc(attachmentName, fileType, attachment.AttachmentType, cpf, attachment.Date, attachment.Observation)
	result, err := coll.InsertOne(ctx, doc)
	tools.CheckError(err)
	fmt.Println(result)

	return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("File %s uploaded successfully ", file.Filename)})

}

func prepareDoc(attachmentName string, fileType string, attachmentType string, cpf string, date string, obs string) bson.D {
	attachment := bson.D{{"tipo", fileType}, {"conteudo", attachmentName}}
	exam := bson.D{{"obs", obs}, {"natureza", attachmentType}, {"anexo", attachment}}
	doc := bson.D{{"cpf", cpf}, {"data", date}, {"documento", exam}}
	return doc
}
