package api

import (
  "Hr-Management-System/utils"
  "github.com/gofiber/fiber/v2"

  "Hr-Management-System/dto"
    "github.com/go-playground/validator/v10"
  "Hr-Management-System/dao"
  )

// @Summary      PUT Employees input: Employees
// @Description  PUT Employees API
// @Tags         PUT Employees - Employees
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Employees false "string collection" 
// @Success      200  {array}   dto.Model_Employees "Status OK"
// @Success      202  {array}   dto.Model_Employees "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateEmployees [PUT]

    func UpdateEmployeesApi(c *fiber.Ctx) error {



    inputObj := dto.Employees{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateEmployees(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}