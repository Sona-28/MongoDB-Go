package services

import (
	"context"
	"fmt"
	"MongoDB-Go/config"
	"MongoDB-Go/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RestaurantContext() *mongo.Collection {
	// client, _ := config.ConnectDataBase()
	return config.GetCollection( "sample_restaurants", "restaurants")
}

func GetRestaurant() ([]*models.Restaurant, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{}
	result, err := RestaurantContext().Find(ctx, filter, options.Find().SetLimit(10))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		// fmt.Println(result)
		var restaurants []*models.Restaurant
		for result.Next(ctx) {
			post := &models.Restaurant{}
			err := result.Decode(post)
			if err != nil {
				return nil, err
			}
			// fmt.Println(products)
			restaurants = append(restaurants, post)
		}
		if err := result.Err(); err != nil {
			return nil, err
		}
		if len(restaurants) == 0 {
			return []*models.Restaurant{}, nil
		}
		return restaurants, nil
	}
}
