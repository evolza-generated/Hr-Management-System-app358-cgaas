package dao

import (
	"Hr-Management-System/dbConfig"
	"Hr-Management-System/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "context"
    "errors"
)

func DB_FindallLaptop () (*[]dto.Laptop, error) {
	var objects []dto.Laptop
	results, err := dbConfig.DATABASE.Collection("Laptops").Find(context.Background(), bson.M{})
	if err != nil {
        if err == mongo.ErrNoDocuments {
        	return nil, nil
        } else {
        	return nil, err
        }
    }
	for results.Next(context.Background()) {
		var object dto.Laptop
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Laptop")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
