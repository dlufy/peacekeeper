package repository

import (
	"fmt"

	"github.com/dlufy/peacekeeper/admin"
	"github.com/dlufy/peacekeeper/database"
	"github.com/dlufy/peacekeeper/models"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertNewProduct(product models.Product) error {
	fmt.Printf("Going to insert %+v", product)
	Client := database.GetClient()
	res, err := Client.Database(admin.GetConfig().Database.DBName).Collection("Products").InsertOne(nil, product)
	fmt.Printf("inserted object is %+v and err %+v", res, err)
	return err
}

func GetProduct(productId int) (product models.Product, err error) {
	filter := bson.D{{"productid", productId}}
	Client := database.GetClient()
	err = Client.Database(admin.GetConfig().Database.DBName).Collection("Products").FindOne(nil, filter).Decode(&product)
	fmt.Printf("err : %+v", err)
	return product, err
}
