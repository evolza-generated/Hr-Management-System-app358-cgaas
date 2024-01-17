package dao

import (
	"Hr-Management-System/dbConfig"
    "Hr-Management-System/dto"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_UpdateMobile-Phone (object *dto.Mobile-Phone)  error {

	result, err := dbConfig.DATABASE.Collection("Mobile-Phones").UpdateOne(context.Background(), bson.M{"objectid": object.ObjectId}, bson.M{"$set": object})
	if err != nil {
		return err
	}
	if result.ModifiedCount < 1 && result.MatchedCount != 1 {
		return errors.New("Specified ID not found!")
	}

	return nil
}