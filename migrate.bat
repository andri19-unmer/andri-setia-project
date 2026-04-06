@echo off
echo Menjalankan migrasi database (UP)...

REM Sesuaikan DB_URL dengan URL Database Anda jika menggunakan host native
set DB_URL="postgres://user:password@localhost:5432/myappdb?sslmode=disable"

REM Atau jalankan ini bila menggunakan migrate murni:
REM migrate -path backend/migrations -database %DB_URL% up

echo Cara Alternatif: Jika ada build make, Anda dapat menggunakan 'make migrate-up'.
make migrate-up
