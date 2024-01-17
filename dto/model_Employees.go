package dto

type Employees struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    EmployeesId string `json:"EmployeesId" validate:"required"`
    FirstName string `json:"FirstName" validate:"required"`
    LastName string `json:"LastName" validate:"required"`
    Email string `json:"Email" validate:"required"`
    PhoneNumber string `json:"PhoneNumber" validate:"required"`
    Salary float64 `json:"Salary" validate:"required"`
    HireDate string `json:"HireDate" validate:"required"`
    Address string `json:"Address" validate:"required"`
    Position string `json:"Position" validate:"required"`
    Department string `json:"Department" validate:"required"`
    }

