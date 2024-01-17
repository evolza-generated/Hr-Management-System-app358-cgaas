package api

import (
  "Hr-Management-System/utils"
  "github.com/gofiber/fiber/v2"

  "Hr-Management-System/dto"
    "github.com/go-playground/validator/v10"
  "Hr-Management-System/dao"
  )

// @Summary      POST Salary input: Salary
// @Description  POST Salary API
// @Tags         POST Salary - Salary
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Salary false "string collection" 
// @Success      200  {array}   dto.Model_Salary "Status OK"
// @Success      202  {array}   dto.Model_Salary "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateSalary [POST]

    func CreateSalaryApi(c *fiber.Ctx) error {



    inputObj := dto.Salary{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_CreateSalary(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}