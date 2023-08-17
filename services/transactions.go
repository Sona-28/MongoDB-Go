package services

import (
	"MongoDB-Go/config"
	"MongoDB-Go/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TransContext()(*mongo.Collection){
	// client,_ := config.ConnectDataBase()
	return config.GetCollection("sample_analytics","transactions")

}

func FindElements()([]*models.Transaction, error){
	ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
	// filter := bson.D{}
	filter := bson.M{"transaction_count": bson.D{{"$gt",85},{"$lt",90}}}
	// filter := bson.M{"$and",["transaction_count": bson.D{{"$gt",100}},"account_id":bson.D{{Key: "$gt",Value: 7000}}]} 
	op := options.Find().SetSort(bson.D{{"transaction_count",1}})
	coll,err := TransContext().Find(ctx,filter,op)
	if err != nil{
		return nil,err
	}
	var transactions []*models.Transaction
	for coll.Next(ctx){
		trans := &models.Transaction{}
		err := coll.Decode(trans)
		if err != nil{
			return nil,err
		}
		transactions = append(transactions, trans)
	}
	if len(transactions) == 0{
		return []*models.Transaction{}, nil
	}
	return transactions, nil
}

func FetchAggregateTransactions(){
	ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
	matchStage := bson.D{{"$match",bson.D{{"account_id",278866}}}}
	groupStage := bson.D{
		{"$group",bson.D{
			{"_id","$account_id"},
			{"total_count",bson.D{{"$sum","$transaction_count"}}},
		}}}
	result,err := TransContext().Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil{
		fmt.Println(err.Error())
	}else{
		var showWithInfo []bson.M
		if err := result.All(ctx,&showWithInfo);err!=nil{
			panic(err)
		}
		// fmt.Println(showWithInfo)
		f_data,err:= json.MarshalIndent(showWithInfo,""," ")
		if err!=nil{
			fmt.Println(err.Error())
		}else{
			fmt.Println(string(f_data))
		}
	}
}

func UpdateTransaction(iv int,nv int) (*mongo.UpdateResult, error){
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	id := bson.D{{"account_id",iv}}
	trans := bson.D{{"$set",bson.D{{"account_id",nv}}}}
	result, err :=  TransContext().UpdateOne(ctx,id,trans)
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}
	// fmt.Println("result",result)
	return result, nil
}
