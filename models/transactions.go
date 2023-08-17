package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Trans struct{
	Date primitive.DateTime `json:"date" bson:"date"`
	Amount int `json:"amount" bson:"amount"`
	TransactionCode string `json:"transaction_code" bson:"transaction_code"`
	Symbol string `json:"symbol" bson:"symbol"`
	Price string `json:"price" bson:"price"`
	Total string `json:"total" bson:"total"`
}

type Transaction struct{
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	AccountId int `json:"account_id" bson:"account_id"`
	TransactionCount int `json:"transaction_count" bson:"transaction_count"`
	BucketStartDate primitive.DateTime `json:"bucket_start_date" bson:"bucket_start_date"`
	BucketEndDate primitive.DateTime `json:"bucket_end_date" bson:"bucket_end_date"`
	// Transactions []Trans `json:"transactions" bson:"transactions"`
}