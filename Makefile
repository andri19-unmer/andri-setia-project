.PHONY: setup run build down logs migrate-up migrate-down

# Variabel koneksi DB (sesuaikan dengan isi .env Anda untuk run native luar docker)
DB_URL="postgres://myuser:mypassword@localhost:5432/myappdb?sslmode=disable"

setup:
	@echo "Menyalin .env.example menjadi .env.dev, .env.staging, dan .env.prod..."
	cp .env.example .env.dev
	cp .env.example .env.staging
	cp .env.example .env.prod
	@echo "Selesai. Silakan periksa file .env.* dan sesuaikan nilainya."

run:
	@echo "Menjalankan aplikasi di environment Development menggunakan Docker..."
	docker-compose --env-file .env.dev up -d

build:
	@echo "Membangun ulang image Docker..."
	docker-compose --env-file .env.dev build

down:
	@echo "Mematikan dan menghapus container..."
	docker-compose down

logs:
	@echo "Menampilkan log aplikasi..."
	docker-compose logs -f

migrate-up:
	@echo "Menjalankan migrasi database (UP)..."
	migrate -path backend/migrations -database $(DB_URL) up

migrate-down:
	@echo "Membatalkan migrasi database (DOWN)..."
	migrate -path backend/migrations -database $(DB_URL) down

clean: down
	@echo "Menghapus volumes database (Hati-Hati!)..."
	docker-compose down -v
