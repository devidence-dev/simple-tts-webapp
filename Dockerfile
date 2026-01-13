FROM python:3.13-slim

WORKDIR /app

# Instalar dependencias del sistema
RUN apt-get update && apt-get install -y --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*

# Copiar pyproject.toml y requirements
COPY pyproject.toml .

# Instalar dependencias Python
RUN pip install --no-cache-dir -e .

# Copiar archivos de la app
COPY main.py .
COPY index.html .

# Exponer puerto
EXPOSE 8000

# Comando para ejecutar la app
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
