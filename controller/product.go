package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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

		imageURL := req.FormValue("imageURL")
		if imageURL != "" {
			res, err := http.Get(imageURL)
			if err != nil {
				log.Println("error while getting image from this URL", imageURL, err)
				return "", err
			}
			tmp := strings.Split(imageURL, "/")
			image := tmp[len(tmp)-1]
			imageList := strings.Split(image, ".")
			if len(imageList) != 2 {
				return "", errors.New(fmt.Sprintf("image path is not correct %s", imageURL))
			}
			return aws.UploadFileToS3(res.Body, admin.GetConfig().AWS.BucketName, strings.ToLower(image), "image/"+imageList[1], "public-read")
		}
		file, header, err := req.FormFile("file")
		if err != nil {
			return "", err
		}
		tempList := strings.Split(strings.ToLower(header.Filename), ".")
		fileExtension := tempList[len(tempList)-1]
		return aws.UploadFileToS3(file, admin.GetConfig().AWS.BucketName, strings.ToLower(header.Filename), "image/"+fileExtension, "public-read")
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
