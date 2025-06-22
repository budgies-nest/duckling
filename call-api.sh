#!/bin/bash

curl http://localhost:7070/api/chat \
-H "Content-Type: application/json" \
-d '
{
  "user": "I need a diagram showing the usage of a JWT in a web application."
}' 