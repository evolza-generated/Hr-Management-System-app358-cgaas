package api

import (
  "Hr-Management-System/utils"
  "github.com/gofiber/fiber/v2"

  "Hr-Management-System/dao"
  )

// @Summary      DELETE Salary input: Salary
// @Description  DELETE Salary API
// @Tags         DELETE Salary - Salary
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Salary "Status OK"
// @Success      202  {array}   dto.Model_Salary "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteSalary [DELETE]

    func DeleteSalaryApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    err := dao.DB_DeleteSalary(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}