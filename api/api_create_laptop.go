package api

import (
  "Hr-Management-System/utils"
  "github.com/gofiber/fiber/v2"

  "Hr-Management-System/dto"
    "github.com/go-playground/validator/v10"
  "Hr-Management-System/dao"
  )

// @Summary      POST Laptop input: Laptop
// @Description  POST Laptop API
// @Tags         POST Laptop - Laptop
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Laptop false "string collection" 
// @Success      200  {array}   dto.Model_Laptop "Status OK"
// @Success      202  {array}   dto.Model_Laptop "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateLaptop [POST]

    func CreateLaptopApi(c *fiber.Ctx) error {



    inputObj := dto.Laptop{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_CreateLaptop(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}