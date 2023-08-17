package services

import (
	"context"
	"fmt"
	"mongodb_connect/config"
	"mongodb_connect/models"
	"time"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
func ProductContext() *mongo.Collection{
	client,_ := config.ConnectDataBase()
	return config.GetCollection(client, "inventory", "products")
}

func InsertProduct(product models.Product) (*mongo.InsertOneResult, error){
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err :=  ProductContext().InsertOne(ctx,product)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("result",result)
	return result, nil
}


func InsertManyProduct(product []interface{}) (*mongo.InsertManyResult, error){
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err :=  ProductContext().InsertMany(ctx,product)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("result",result)
	return result, nil
}


func UpdateProduct(id interface{},product interface{}) (*mongo.UpdateResult, error){
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err :=  ProductContext().UpdateOne(ctx,id, product)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("result",result)
	return result, nil
}

func GetProduct() ([]*models.Product,error) {
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{Key: "price",Value: 115000}}
	result, err :=  ProductContext().Find(ctx,filter)
	if err!=nil{
		fmt.Println(err.Error())
		return nil,err
	}else{
		fmt.Println(result)
		var products []*models.Product
		for result.Next(ctx){
			post := &models.Product{}
			err := result.Decode(post)
			if err!=nil{
				return nil,err
			}
			// fmt.Println(products)
			products = append(products, post)
		}
		if err:=result.Err(); err!=nil{
			return nil, err
		}
		if len(products) == 0{
			return []*models.Product{},nil
		}
		return products,nil
	}
}