package main

import (
	"log"
	"strconv"
	"fmt"
	"os"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	app.Get("/retoibm/sumar/:sum01/:sum02", func(c *fiber.Ctx) error {
		sum01,err1 := strconv.ParseInt(c.Params("sum01"), 10, 0)
		sum02,err2 := strconv.ParseInt(c.Params("sum02"), 10, 0)
		hostname, err3 := os.Hostname()

		_ = err1
		_ = err2
		_ = err3

        sum_result := sum01 + sum02
		result := fmt.Sprintf(`{ "resultado" : %s, "hostname": "%s" , "language": "%s" }`, strconv.Itoa(int(sum_result)), hostname,"go")

        return c.SendString(result)
    })

	// Start server
	log.Fatal(app.Listen(":3000"))
}
