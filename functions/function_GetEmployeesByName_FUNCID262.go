package functions

import (
	"context"
	"Hr-Management-System/dto"
	"Hr-Management-System/dbConfig"
	"go.mongodb.org/mongo-driver/bson"
)

func GetEmployeesByName_FUNCID262(name string) ([]dto.Employees, error) {
	employees := []dto.Employees{}
	
	filter := bson.M{"FirstName": name}
	
	cursor, err := dbConfig.DATABASE.Collection("Employees").Find(context.Background(), filter)
	if err != nil {
		return employees, err
	}
	defer cursor.Close(context.Background())
	
	for cursor.Next(context.Background()) {
		var emp dto.Employees
		err := cursor.Decode(&emp)
		if err != nil {
			return employees, err
		}
		employees = append(employees, emp)
	}
	
	if err := cursor.Err(); err != nil {
		return employees, err
	}
	
	return employees, nil
}