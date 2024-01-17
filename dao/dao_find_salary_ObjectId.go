package dao

import (
	"Hr-Management-System/dbConfig"
	"Hr-Management-System/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindSalarybyObjectId (objectid string) (*dto.Salary, error) {
	var object dto.Salary

	err := dbConfig.DATABASE.Collection("Salarys").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
