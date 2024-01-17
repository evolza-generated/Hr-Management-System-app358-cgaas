package dao

import (
	"Hr-Management-System/dbConfig"
	"Hr-Management-System/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "context"
    "errors"
)

func DB_FindallMobile-Phone () (*[]dto.Mobile-Phone, error) {
	var objects []dto.Mobile-Phone
	results, err := dbConfig.DATABASE.Collection("Mobile-Phones").Find(context.Background(), bson.M{})
	if err != nil {
        if err == mongo.ErrNoDocuments {
        	return nil, nil
        } else {
        	return nil, err
        }
    }
	for results.Next(context.Background()) {
		var object dto.Mobile-Phone
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Mobile-Phone")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
