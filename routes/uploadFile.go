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
	"mime/multipart"
	"net/http"
	"time"
)

type Attachment struct {
	AttachmentType   string                `json:"attachmentType"`
	AdditionalFields map[string]string     `json:"additionalFields"`
	File             *multipart.FileHeader `json:"file"`
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
	// TODO: Descobrir por que não está vindo os parâmetros na requisição
	if err := c.Bind(attachment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//attachment.AttachmentType = "exame"
	//attachment.AdditionalFields = map[string]string{"tipo": "Teste Overdose"}

	file, err := c.FormFile("File")
	tools.CheckError(err)
	filename := file.Filename
	fileContent, err := file.Open()
	t := time.Now().Format("20060102150405")
	attachmentName := t + filename
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
	doc := prepareDoc(attachmentName, fileType, attachment.AttachmentType, cpf, attachment.AdditionalFields)
	result, err := coll.InsertOne(ctx, doc)
	tools.CheckError(err)
	fmt.Println(result)

	return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("File %s uploaded successfully ", file.Filename)})

}

func prepareDoc(attachmentName string, fileType string, attachmentType string, cpf string, additionalFields map[string]string) bson.D {
	if attachmentType == "exame" {
		attachment := bson.D{{"tipo", fileType}, {"conteudo", attachmentName}}
		exam := bson.D{{"tipo", additionalFields["tipo"]}, {"anexo", attachment}}
		doc := bson.D{{"cpf", cpf}, {"exame", exam}}

		return doc
	}

	panic("Invalid attachment type")
}
