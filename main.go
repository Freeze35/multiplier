package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"log"
	handler "multiplier/internal"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {

	// Read string rtp flag
	var nFlag = flag.String("rtp", "1234.0", "help message for flag n")
	flag.Parse()

	// Converting string rtp flag to float64
	rtp, err := strconv.ParseFloat(*nFlag, 64)
	if err != nil {
		log.Fatalf("invalid value: %v", err)
	}

	// Check rtp number
	if rtp <= 0 || rtp > 1.0 {
		log.Fatalf("invalid rtp: %f", rtp)
	}

	// Init NewHandler
	h, err := handler.NewHandler(rtp) // Предполагается, что NewHandler принимает rtp
	if err != nil {
		log.Fatal("An initialization error of the handler: ", err)
	}

	app := fiber.New()
	h.InitRoutes(app)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start server
	go func() {
		if err := app.Listen(":64333"); err != nil {
			log.Fatal("Server failed to start: ", err)
		}
	}()

	<-quit

	// Shutdown server
	if err := app.Shutdown(); err != nil {
		log.Fatal("Server shutdown failed: ", err)
	}

}
