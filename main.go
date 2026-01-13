package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

type TTSReq struct {
	Text   string  `json:"text"`
	Voice  string  `json:"voice"`
	Format string  `json:"format"`
	Speed  float64 `json:"speed"`
}

func main() {
	kokoroBase := os.Getenv("KOKORO_BASE")
	if kokoroBase == "" {
		kokoroBase = "http://localhost:8880"
	}

	client := &http.Client{}

	app := fiber.New()

	// Endpoint to get voices
	app.Get("/api/voices", func(c *fiber.Ctx) error {
		resp, err := client.Get(kokoroBase + "/v1/audio/voices")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		c.Set("Content-Type", "application/json")
		return c.Send(body)
	})

	// Endpoint for TTS
	app.Post("/api/tts", func(c *fiber.Ctx) error {
		var req TTSReq
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
		}

		// Set defaults
		if req.Voice == "" {
			req.Voice = "af_bella"
		}
		if req.Format == "" {
			req.Format = "mp3"
		}
		if req.Speed == 0 {
			req.Speed = 1.0
		}

		payload := map[string]interface{}{
			"model":           "kokoro",
			"input":           req.Text,
			"voice":           req.Voice,
			"response_format": req.Format,
			"speed":           req.Speed,
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		resp, err := client.Post(kokoroBase+"/v1/audio/speech", "application/json", bytes.NewReader(jsonPayload))
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		defer resp.Body.Close()

		mediaTypes := map[string]string{
			"mp3":  "audio/mpeg",
			"wav":  "audio/wav",
			"opus": "audio/ogg",
			"flac": "audio/flac",
			"m4a":  "audio/mp4",
			"pcm":  "application/octet-stream",
		}
		contentType := mediaTypes[req.Format]
		if contentType == "" {
			contentType = "application/octet-stream"
		}

		c.Set("Content-Type", contentType)
		return c.SendStream(resp.Body)
	})

	// Serve static files
	app.Static("/", "./")

	log.Fatal(app.Listen(":8000"))
}
