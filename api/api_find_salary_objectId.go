package api

import (
  "Hr-Management-System/utils"
  "github.com/gofiber/fiber/v2"

  "Hr-Management-System/dao"
  )

// @Summary      GET Salary input: Salary
// @Description  GET Salary API
// @Tags         GET Salary - Salary
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Salary "Status OK"
// @Success      202  {array}   dto.Model_Salary "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindSalary [GET]

    func FindSalaryApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    returnValue, err := dao.DB_FindSalarybyObjectId(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}