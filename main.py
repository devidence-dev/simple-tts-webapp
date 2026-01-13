from fastapi import FastAPI
from fastapi.responses import StreamingResponse, FileResponse
from fastapi.staticfiles import StaticFiles
import httpx
from pydantic import BaseModel
from pathlib import Path
import os

KOKORO_BASE = os.getenv("KOKORO_BASE", "http://localhost:8880")

app = FastAPI()

class TTSReq(BaseModel):
    text: str
    voice: str = "af_bella"
    format: str = "mp3"   # mp3, wav, opus, flac, m4a, pcm
    speed: float = 1.0

@app.get("/api/voices")
async def voices():
    async with httpx.AsyncClient() as client:
        r = await client.get(f"{KOKORO_BASE}/v1/audio/voices")
        r.raise_for_status()
        return r.json()

@app.post("/api/tts")
async def tts(req: TTSReq):
    payload = {
        "model": "kokoro",
        "input": req.text,
        "voice": req.voice,
        "response_format": req.format,
        "speed": req.speed,
    }
    async with httpx.AsyncClient(timeout=None) as client:
        r = await client.post(f"{KOKORO_BASE}/v1/audio/speech", json=payload)
        r.raise_for_status()

    media = {
        "mp3": "audio/mpeg",
        "wav": "audio/wav",
        "opus": "audio/ogg",
        "flac": "audio/flac",
        "m4a": "audio/mp4",
        "pcm": "application/octet-stream",
    }.get(req.format, "application/octet-stream")

    return StreamingResponse(iter([r.content]), media_type=media)

# Serve static files at the end
static_dir = Path(__file__).parent
app.mount("/", StaticFiles(directory=static_dir, html=True), name="static")
