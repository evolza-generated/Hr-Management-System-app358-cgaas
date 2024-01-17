package functions

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"Hr-Management-System/dto"
	"Hr-Management-System/dbConfig"
)

func GetEmployeesByFirstName_FUNCID261(firstName string) ([]dto.Employees, error) {
	collection := dbConfig.DATABASE.Collection("Employees")
	filter := bson.M{"FirstName": firstName}
	var employees []dto.Employees

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var emp dto.Employees
		if err := cursor.Decode(&emp); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}