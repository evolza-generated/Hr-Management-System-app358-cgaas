package api

import (
  "Hr-Management-System/utils"
  "github.com/gofiber/fiber/v2"

  "Hr-Management-System/dao"
  )

// @Summary      GET Employees input: Employees
// @Description  GET Employees API
// @Tags         GET Employees - Employees
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Employees "Status OK"
// @Success      202  {array}   dto.Model_Employees "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindallEmployees [GET]

    func FindallEmployeesApi(c *fiber.Ctx) error {


returnValue, err := dao.DB_FindallEmployees()
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}