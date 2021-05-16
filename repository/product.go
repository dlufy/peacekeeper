package repository

import (
	"fmt"
	"log"

	"github.com/dlufy/peacekeeper/admin"
	"github.com/dlufy/peacekeeper/database"
	"github.com/dlufy/peacekeeper/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertNewProduct(product models.Product) error {
	product.ProductID = primitive.NewObjectID()
	fmt.Printf("Going to insert %+v", product)
	Client := database.GetClient()
	res, err := Client.Database(admin.GetConfig().Database.DBName).Collection("Products").InsertOne(nil, product)
	fmt.Printf("inserted object is %+v and err %+v", res, err)
	return err
}

func GetProduct(productId string) (product models.Product, err error) {

	log.Println("[GetProduct]info ID", productId)
	filter := GetFilterMap(productId, "", 0, 1000)
	Client := database.GetClient()
	err = Client.Database(admin.GetConfig().Database.DBName).Collection("Products").FindOne(nil, filter).Decode(&product)
	if err != nil {
		fmt.Printf("err : %+v", err)
	}

	return product, err
}

func GetFilterMap(productId, name string, PriceMin, PriceMax int64) bson.M {

	docID, err := primitive.ObjectIDFromHex(productId)
	filter := bson.M{}
	if err != nil {
		log.Println("not able to convert productId to objectId", err)
	} else {
		filter["_id"] = docID
	}
	if PriceMax == 0 {
		PriceMax = 100000
	}
	return filter
}
