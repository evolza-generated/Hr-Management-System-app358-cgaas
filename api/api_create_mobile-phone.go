package api

import (
  "Hr-Management-System/utils"
  "github.com/gofiber/fiber/v2"

  "Hr-Management-System/dto"
    "github.com/go-playground/validator/v10"
  "Hr-Management-System/dao"
  )

// @Summary      POST Mobile-Phone input: Mobile-Phone
// @Description  POST Mobile-Phone API
// @Tags         POST Mobile-Phone - Mobile-Phone
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Mobile-Phone false "string collection" 
// @Success      200  {array}   dto.Model_Mobile-Phone "Status OK"
// @Success      202  {array}   dto.Model_Mobile-Phone "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateMobile-Phone [POST]

    func CreateMobile-PhoneApi(c *fiber.Ctx) error {



    inputObj := dto.Mobile-Phone{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_CreateMobile-Phone(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}