#!/bin/bash
echo "Menjalankan container Docker..."
docker-compose --env-file .env.dev up -d
echo "Container berjalan di background."
