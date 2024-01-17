package dao

import (
    "context"
	"Hr-Management-System/dbConfig"
	"Hr-Management-System/dto"

)

func DB_CreateSalary (application *dto.Salary) error {

	_, err := dbConfig.DATABASE.Collection("Salarys").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}