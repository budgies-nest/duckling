#!/bin/bash

curl http://localhost:7070/api/chat \
-H "Content-Type: application/json" \
-d '
{
  "user": "who is James T Kirk?"
}' 