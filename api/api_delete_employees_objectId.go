package api

import (
  "Hr-Management-System/utils"
  "github.com/gofiber/fiber/v2"

  "Hr-Management-System/dao"
  )

// @Summary      DELETE Employees input: Employees
// @Description  DELETE Employees API
// @Tags         DELETE Employees - Employees
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Employees "Status OK"
// @Success      202  {array}   dto.Model_Employees "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteEmployees [DELETE]

    func DeleteEmployeesApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    err := dao.DB_DeleteEmployees(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}