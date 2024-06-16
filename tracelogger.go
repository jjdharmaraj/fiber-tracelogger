package tracelogger

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

// Logger middleware logs the request when the logger level is set to trace.
func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if os.Getenv("LOGGER_LEVEL") == "trace" {
			// Extract relevant request details
			method := c.Method()
			path := c.Path()
			ip := c.IP()
			userAgent := c.Get("User-Agent")

			log.Printf("method: %s, path: %s, ip: %s, user-agent: %s\n", method, path, ip, userAgent)

			// Include request body for POST/PUT requests (optional)
			if c.Method() == http.MethodPost || c.Method() == http.MethodPut {
				body := c.Body()
				if len(body) == 0 {
					log.Printf("Error reading request body: %v\n", body)
				} else {
					//This wont work for other sorts of body inputs.
					var data map[string]interface{}
					if err := json.Unmarshal(body, &data); err != nil {
						return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
					} else {
						log.Printf("request body: %s\n", data) // Use only the body
					}
				}

			}
		}
		return c.Next()
	}
}
