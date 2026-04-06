@echo off
echo Membangun ulang Docker image menggunakan .env.dev...
docker-compose --env-file .env.dev build
echo Selesai membangun image.
