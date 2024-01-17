package apiHandlers

import (
	"Hr-Management-System/api"

	"encoding/gob"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Router(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	store := session.New()
	//encryptcookie.New(encryptcookie.Config{
	//	Key: "secret-thirty-2-character-string",
	//})
	gob.Register(map[string]interface{}{})

	api := app.Group("/Hr-Management-System/api")
	check := app.Group("/")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	DefaultMappings(check, store)
	RouteMappings(api,store)
}

func RouteMappings(cg fiber.Router, store *session.Store) {
cg.Post("/CreateEmployees", api.CreateEmployeesApi)
cg.Put("/UpdateEmployees", api.UpdateEmployeesApi)
cg.Delete("/DeleteEmployees", api.DeleteEmployeesApi)
cg.Get("/FindEmployees", api.FindEmployeesApi)
cg.Get("/FindallEmployees", api.FindallEmployeesApi)
cg.Post("/CreateSalary", api.CreateSalaryApi)
cg.Put("/UpdateSalary", api.UpdateSalaryApi)
cg.Delete("/DeleteSalary", api.DeleteSalaryApi)
cg.Get("/FindSalary", api.FindSalaryApi)
cg.Get("/FindallSalary", api.FindallSalaryApi)
cg.Post("/CreateLaptop", api.CreateLaptopApi)
cg.Put("/UpdateLaptop", api.UpdateLaptopApi)
cg.Delete("/DeleteLaptop", api.DeleteLaptopApi)
cg.Get("/FindLaptop", api.FindLaptopApi)
cg.Get("/FindallLaptop", api.FindallLaptopApi)
cg.Post("/CreateMobilePhone", api.CreateMobilePhoneApi)
cg.Put("/UpdateMobilePhone", api.UpdateMobilePhoneApi)
cg.Delete("/DeleteMobilePhone", api.DeleteMobilePhoneApi)
cg.Get("/FindMobilePhone", api.FindMobilePhoneApi)
cg.Get("/FindallMobilePhone", api.FindallMobilePhoneApi)
cg.Get("/GetEmployeesByFirstName_FUNCID261", api.GetEmployeesByFirstName_FUNCID261Api)
cg.Get("/GetEmployeesByName_FUNCID262", api.GetEmployeesByName_FUNCID262Api)


}

func DefaultMappings(cg fiber.Router, store *session.Store) {
	cg.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "Hr-Management-System-APP358 service is up and running", "version": "1.0"})
	})
}