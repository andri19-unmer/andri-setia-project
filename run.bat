@echo off
echo Menjalankan aplikasi dengan Docker Compose di background...
docker-compose --env-file .env.dev up -d
echo Aplikasi berjalan. Gunakan "docker-compose logs -f" untuk melihat log.
