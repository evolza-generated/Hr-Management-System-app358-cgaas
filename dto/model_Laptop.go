package dto

type Laptop struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    }

