#!/bin/bash

curl --no-buffer http://localhost:7070/api/chat-stream \
-H "Content-Type: application/json" \
-d '
{
  "user": "who is Jean-Luc Picard?"
}' 
