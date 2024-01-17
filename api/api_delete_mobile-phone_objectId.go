package api

import (
  "Hr-Management-System/utils"
  "github.com/gofiber/fiber/v2"

  "Hr-Management-System/dao"
  )

// @Summary      DELETE Mobile-Phone input: Mobile-Phone
// @Description  DELETE Mobile-Phone API
// @Tags         DELETE Mobile-Phone - Mobile-Phone
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Mobile-Phone "Status OK"
// @Success      202  {array}   dto.Model_Mobile-Phone "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteMobile-Phone [DELETE]

    func DeleteMobile-PhoneApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    err := dao.DB_DeleteMobile-Phone(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}