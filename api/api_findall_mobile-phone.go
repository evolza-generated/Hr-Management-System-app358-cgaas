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
// @Router      /FindallMobile-Phone [GET]

    func FindallMobile-PhoneApi(c *fiber.Ctx) error {


returnValue, err := dao.DB_FindallMobile-Phone()
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}