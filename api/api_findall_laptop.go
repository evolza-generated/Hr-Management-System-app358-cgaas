package api

import (
  "Hr-Management-System/utils"
  "github.com/gofiber/fiber/v2"

  "Hr-Management-System/dao"
  )

// @Summary      GET Laptop input: Laptop
// @Description  GET Laptop API
// @Tags         GET Laptop - Laptop
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Laptop "Status OK"
// @Success      202  {array}   dto.Model_Laptop "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindallLaptop [GET]

    func FindallLaptopApi(c *fiber.Ctx) error {


returnValue, err := dao.DB_FindallLaptop()
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}