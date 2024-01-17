package dao

import (
	"Hr-Management-System/dbConfig"
    "Hr-Management-System/dto"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_UpdateLaptop (object *dto.Laptop)  error {

	result, err := dbConfig.DATABASE.Collection("Laptops").UpdateOne(context.Background(), bson.M{"objectid": object.ObjectId}, bson.M{"$set": object})
	if err != nil {
		return err
	}
	if result.ModifiedCount < 1 && result.MatchedCount != 1 {
		return errors.New("Specified ID not found!")
	}

	return nil
}