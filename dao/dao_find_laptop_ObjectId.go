package dao

import (
	"Hr-Management-System/dbConfig"
	"Hr-Management-System/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindLaptopbyObjectId (objectid string) (*dto.Laptop, error) {
	var object dto.Laptop

	err := dbConfig.DATABASE.Collection("Laptops").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
