package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dlufy/peacekeeper/admin"
	core "github.com/dlufy/peacekeeper/admin"
	"github.com/dlufy/peacekeeper/aws"
	"github.com/dlufy/peacekeeper/models"
	"github.com/dlufy/peacekeeper/repository"
)

func HandleUploadRequest(res http.ResponseWriter, req *http.Request) (data interface{}, err error) {
	switch strings.ToUpper(req.Method) {
	case "GET":
		fileURL := req.URL.Query().Get("fileURL")
		fileName := "tempfile.png"
		_, err = aws.DownloadFileFromS3(fileURL, fileName, "", core.GetConfig().AWS.BucketName)
		return "", err
	case "POST":
		file, header, err := req.FormFile("file")
		tempList := strings.Split(strings.ToLower(header.Filename), ".")
		fileExtension := tempList[len(tempList)-1]
		url, err := aws.UploadFileToS3(file, admin.GetConfig().AWS.BucketName, strings.ToLower(header.Filename), "image/"+fileExtension, "public-read")
		return url, err
	}
	return "", nil
}
func HandleProductRequest(res http.ResponseWriter, req *http.Request) (data interface{}, err error) {
	switch strings.ToUpper(req.Method) {
	case "GET":
		productId := req.URL.Query().Get("product_id")
		return repository.GetProduct(productId)
	case "POST":
		byt, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		var Product models.Product
		err = json.Unmarshal(byt, &Product)
		if err != nil {
			fmt.Println("error while unmarshalling ", string(byt), err)
			return nil, err
		}
		err = repository.InsertNewProduct(Product)
		data = Product
	default:
		res.WriteHeader(http.StatusOK)
	}
	return nil, err
}
