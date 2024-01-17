package api

import (
  "Hr-Management-System/utils"
  "github.com/gofiber/fiber/v2"

  "Hr-Management-System/dto"
    "github.com/go-playground/validator/v10"
  "Hr-Management-System/dao"
  )

// @Summary      PUT Mobile-Phone input: Mobile-Phone
// @Description  PUT Mobile-Phone API
// @Tags         PUT Mobile-Phone - Mobile-Phone
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Mobile-Phone false "string collection" 
// @Success      200  {array}   dto.Model_Mobile-Phone "Status OK"
// @Success      202  {array}   dto.Model_Mobile-Phone "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateMobile-Phone [PUT]

    func UpdateMobile-PhoneApi(c *fiber.Ctx) error {



    inputObj := dto.Mobile-Phone{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateMobile-Phone(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}