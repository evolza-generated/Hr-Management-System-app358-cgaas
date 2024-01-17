package dao

import (
    "context"
	"Hr-Management-System/dbConfig"
	"Hr-Management-System/dto"

)

func DB_CreateMobile-Phone (application *dto.Mobile-Phone) error {

	_, err := dbConfig.DATABASE.Collection("Mobile-Phones").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}