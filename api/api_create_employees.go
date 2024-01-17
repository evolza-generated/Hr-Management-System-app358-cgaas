package api

import (
  "Hr-Management-System/utils"
  "github.com/gofiber/fiber/v2"

  "Hr-Management-System/functions"
    "Hr-Management-System/dto"
    "github.com/go-playground/validator/v10"
  "Hr-Management-System/dao"
  )

// @Summary      POST Employees input: Employees
// @Description  POST Employees API
// @Tags         POST Employees - Employees
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Employees false "string collection" 
// @Success      200  {array}   dto.Model_Employees "Status OK"
// @Success      202  {array}   dto.Model_Employees "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateEmployees [POST]

    func CreateEmployeesApi(c *fiber.Ctx) error {



    inputObj := dto.Employees{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    
if err := functions.UniqueCheck(inputObj, "Employeess", []string{ "EmployeesId", "Email"}) ; err!=nil{
          return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
        }
    
    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_CreateEmployees(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}