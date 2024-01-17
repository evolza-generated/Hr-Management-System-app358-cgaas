package api

import (
  "Hr-Management-System/utils"
  "github.com/gofiber/fiber/v2"

  "Hr-Management-System/dao"
  )

// @Summary      DELETE Laptop input: Laptop
// @Description  DELETE Laptop API
// @Tags         DELETE Laptop - Laptop
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Laptop "Status OK"
// @Success      202  {array}   dto.Model_Laptop "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteLaptop [DELETE]

    func DeleteLaptopApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    err := dao.DB_DeleteLaptop(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}