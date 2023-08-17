package main

import (
	// "context"
	"fmt"
	// "mongodb_connect/config"
	// "mongodb_connect/models"
	"mongodb_connect/services"

	// "golang.org/x/vuln/client"
	// "mongodb_connect/constants"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

// var (
// 	mongoCient *mongo.Client
// )

func main() {
	// client, err := config.ConnectDataBase()
	// config.GetCollection(client,sample_training)
	// product := models.Product{ID: primitive.NewObjectID(),
	// 	Name:        "Iphone",
	// 	Price:       151000,
	// 	Description: "Budget phone"}
	// services.InsertProduct(product)

	// p1 := []interface{}{models.Product{ID: primitive.NewObjectID(),
	// 	Name:        "Iphone",
	// 	Price:       151000,
	// 	Description: "Nice phone"},
	// 	models.Product{ID: primitive.NewObjectID(),
	// 		Name:        "Samsung",
	// 		Price:       50000,
	// 		Description: "Budget phone"},}
	// services.InsertManyProduct(p1)

	// id,_ := primitive.ObjectIDFromHex("64dc98cd88ffadd9376df516")
	// filter := bson.D{{Key: "_id",Value: id}}
	// update := bson.D{{Key: "$set",Value: bson.D{{Key: "rating",Value: 9.2}}}}
	// services.UpdateProduct(filter,update)

	
	// products,_ := services.GetProduct()
	// for _,pro := range products{
	// 	fmt.Println(pro)
	// }

	r,_ := services.GetRestaurant()
	for _,res := range r{
		fmt.Println(res.Name)
	}
	fmt.Println("MongoDB connected successfully...")
}
