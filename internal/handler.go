package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
	"time"
)

type Handler struct {
	rtp              float64
	clientSum        float64 // The total amount based on transformation
	numberIterations float64 // To check the quantity, the client's value
}

func NewHandler(rtp float64) (*Handler, error) {

	//Initializes the random number generator with a unique seed value
	rand.Seed(time.Now().UnixNano())

	return &Handler{rtp: rtp, clientSum: 0, numberIterations: 0}, nil
}

const (
	limitSizeNumber = 10000.0
)

func (h *Handler) InitRoutes(app *fiber.App) {
	app.Get("/get", func(c *fiber.Ctx) error {

		x := c.QueryFloat("x")
		if x < 1.0 || x > 10000.0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "'x' must be in [1.0, 10000.0]",
			})
		}

		// The generation of the multiplier
		multiplier := h.generateMultiplier(x)

		return c.JSON(fiber.Map{
			"result": fmt.Sprintf("%.1f", multiplier),
		})
	})
}

func (h *Handler) generateMultiplier(x float64) (multiplier float64) {

	//Incrementing client's numbers
	h.numberIterations += 1

	generatedNumber := h.generateNumber(x)
	multiplier = generatedNumber

	//Checking returning coefficient
	log.Printf("%.1f It,%v", h.clientSum/h.numberIterations, h.numberIterations)

	return multiplier
}

func (h *Handler) generateNumber(x float64) float64 {

	//We check if our coefficient is more than rtp or not
	if h.clientSum/h.numberIterations > h.rtp && h.clientSum != 0 {
		newNumber := 1.0 + rand.Float64()*(x-1.0)
		return newNumber
	} else {
		newNumber := 1.0 + rand.Float64()*(limitSizeNumber-1.0)
		//We increase the total amount of the client
		h.clientSum += x
		return newNumber
	}
}
