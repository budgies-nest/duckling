# 🦆 Duckling
> Quick AI Agent with Docker Compose

**Compose file**:
```yaml
services:
  duckling-agent:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - 7070:${HTTP_PORT}
    environment:
      - HTTP_PORT=${HTTP_PORT}
      - ENGINE_ENDPOINT=${DMR_BASE_URL}/engines/llama.cpp/v1
      - DMR_CHAT_MODEL=${DMR_CHAT_MODEL}
      - SYSTEM_INSTRUCTION=${SYSTEM_INSTRUCTION}
    depends_on:
      - download-chat-model

  download-chat-model:
    provider:
      type: model
      options:
        model: ${DMR_CHAT_MODEL}

```

**Start the Duckling AI Agent**:

```bash
docker compose up -d
```

**Test the Duckling AI Agent**:

```bash
curl --no-buffer http://localhost:7070/api/chat-stream \
-H "Content-Type: application/json" \
-d '
{
  "user": "I need a diagram showing the usage of a JWT in a web application."
}' 
```
