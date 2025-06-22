# 🦆 Duckling
> Quick AI Agents with Docker Compose


## Start the Duckling AI Agents

```bash
docker compose up -d
```

## Test the Duckling AI Agents

**Mermaid Agent**:
```bash
curl --no-buffer http://localhost:7071/api/chat-stream \
-H "Content-Type: application/json" \
-d '
{
  "user": "I need a diagram showing the usage of a JWT in a web application."
}' 
```

**Golang Agent**:
```bash
curl --no-buffer http://localhost:7072/api/chat-stream \
-H "Content-Type: application/json" \
-d '
{
  "user": "I need a function to write text to file and another function read it back."
}' 
```