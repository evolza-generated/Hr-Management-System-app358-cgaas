package api

import (
  "Hr-Management-System/utils"
  "github.com/gofiber/fiber/v2"

  "Hr-Management-System/dao"
  )

// @Summary      GET Mobile-Phone input: Mobile-Phone
// @Description  GET Mobile-Phone API
// @Tags         GET Mobile-Phone - Mobile-Phone
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Mobile-Phone "Status OK"
// @Success      202  {array}   dto.Model_Mobile-Phone "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindMobile-Phone [GET]

    func FindMobile-PhoneApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    returnValue, err := dao.DB_FindMobile-PhonebyObjectId(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}