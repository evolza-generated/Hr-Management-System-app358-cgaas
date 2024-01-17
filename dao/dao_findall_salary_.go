package dao

import (
	"Hr-Management-System/dbConfig"
	"Hr-Management-System/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "context"
    "errors"
)

func DB_FindallSalary () (*[]dto.Salary, error) {
	var objects []dto.Salary
	results, err := dbConfig.DATABASE.Collection("Salarys").Find(context.Background(), bson.M{})
	if err != nil {
        if err == mongo.ErrNoDocuments {
        	return nil, nil
        } else {
        	return nil, err
        }
    }
	for results.Next(context.Background()) {
		var object dto.Salary
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Salary")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
