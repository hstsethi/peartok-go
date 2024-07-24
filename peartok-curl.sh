#!/bin/sh

curl  -v -H "Content-Type: application/json" -d "{\"port\": 9812}" http://localhost:8080/receive

curl  -X   -H "Content-Type: application/json" -d "{\"Amount\":10, \"Receiver\" : \"localhost:9812\"}" http://localhost:8080/send
