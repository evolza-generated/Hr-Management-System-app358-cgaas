package dao

import (
	"Hr-Management-System/dbConfig"
	"Hr-Management-System/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindMobile-PhonebyObjectId (objectid string) (*dto.Mobile-Phone, error) {
	var object dto.Mobile-Phone

	err := dbConfig.DATABASE.Collection("Mobile-Phones").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
