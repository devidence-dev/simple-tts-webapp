# TTS WebApp

A simple web application to convert text to audio using the Kokoro TTS service.

## Description

This application provides an intuitive web interface to generate audio from text using the Kokoro voice synthesis engine. It is built with Fiber for the backend and pure HTML/CSS/JavaScript for the frontend.

## Features

- 🎙️ AI-powered text-to-speech conversion
- 🌐 Responsive web interface
- 🎵 Multiple supported audio formats (MP3, WAV, Opus, FLAC, M4A)
- 🗣️ Various available voices
- ⚡ Playback speed control
- 📥 Direct download of generated audio
- 🐳 Docker deployment

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose (optional, for full deployment)

## Installation

### Option 1: Local Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-user/tts-webapp.git
   cd tts-webapp
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build the application:
   ```bash
   go build -o main .
   ```

4. Ensure the Kokoro service is running on `http://localhost:8880` (see Docker section).

5. Run the application:
   ```bash
   ./main
   ```

6. Open your browser at `http://localhost:8000`

### Option 2: Docker Deployment

1. Clone the repository:
   ```bash
   git clone https://github.com/your-user/tts-webapp.git
   cd tts-webapp
   ```

2. Run with Docker Compose:
   ```bash
   docker-compose up -d
   ```

3. The application will be available at `http://localhost:8000`

## Usage

1. Open the application in your browser
2. Write or paste the text you want to convert in the text area
3. Select the desired voice
4. Choose the audio format
5. Adjust the speed if necessary
6. Click "Generate Audio"
7. Play the generated audio or download it

## API

The application exposes the following endpoints:

### GET /api/voices
Gets the list of available voices.

**Response:**
```json
{
  "voices": ["af_bella", "af_sarah", ...]
}
```

### POST /api/tts
Generates audio from text.

**Request Body:**
```json
{
  "text": "Text to convert",
  "voice": "af_bella",
  "format": "mp3",
  "speed": 1.0
}
```

**Response:** Audio in the specified format

## Dependencies

- Fiber: Web framework for Go
- net/http: Standard HTTP client
- Kokoro TTS: Voice synthesis service (requires separate deployment)

## Contribution

Contributions are welcome. Please open an issue or send a pull request.

## License

This project is under the MIT License.