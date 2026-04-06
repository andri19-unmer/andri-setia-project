#!/bin/bash
echo "Membangun ulang Docker image..."
docker-compose --env-file .env.dev build
echo "Selesai."
