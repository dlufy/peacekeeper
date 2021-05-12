package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/dlufy/peacekeeper/models"
	"github.com/dlufy/peacekeeper/repository"
)

func HandleProductRequest(res http.ResponseWriter, req *http.Request) (data interface{}, err error) {
	switch strings.ToUpper(req.Method) {
	case "GET":
		productId, _ := strconv.Atoi(req.URL.Query().Get("product_id"))
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
	}
	return nil, err
}
