package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Addresss struct {
	Building string    `json:"building" bson:"building"`
	Coord    []float64 `json:"coord" bson:"coord"`
	Street   string    `json:"street" bson:"street"`
	Zipcode  string    `json:"zipcode" bson:"zipcode"`
}

type Grad struct {
	Date  primitive.DateTime `json:"date" bson:"date"`
	Grade string             `json:"grade" bson:"grade"`
	Score int                `json:"score" bson:"score"`
}

type Restaurant struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Address       Addresss           `json:"address" bson:"address"`
	Borough       string             `json:"borough" bson:"borough"`
	Cuisine       string             `json:"cuisine" bson:"cuisine"`
	Grades        []Grad             `json: "grades" bson:"grades"`
	Name          string             `json:"name" bson:"name"`
	Restaurant_id string             `json:"restaurant_id" bson:"restaurant_id"`
}
