package api

import (
  "Hr-Management-System/utils"
  "github.com/gofiber/fiber/v2"

  
    "Hr-Management-System/functions"
  )

// @Summary      CUSTOMG  input: 
// @Description  CUSTOMG  API
// @Tags         CUSTOMG  - 
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_err "Status OK"
// @Success      202  {array}   dto.Model_err "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /Customg [CUSTOMG]

    func GetEmployeesByFirstName_FUNCID261Api(c *fiber.Ctx) error {


firstName := c.Query("firstName")
     
   response ,    err  :=functions.GetEmployeesByFirstName_FUNCID261(firstName)
   
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    




    
    
    return c.Status(fiber.StatusCreated).JSON(response)
    
    
}