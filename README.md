# Full-Stack Go & Vite Web Application

Proyek ini adalah contoh boilerplate/panduan untuk membangun aplikasi web *full-stack* yang siap digunakan di tahap *production*. Aplikasi ini menggunakan backend Golang (Go), frontend React dengan Vite dan TypeScript, serta PostgreSQL sebagai basis data.

## Teknologi Utama

- **Backend**: Golang, Fiber/Echo (Clean Architecture), GORM/sqlc, golang-migrate
- **Frontend**: React, Vite, TypeScript, Tailwind CSS, Zustand, TanStack Query
- **Database**: PostgreSQL
- **Infrastruktur**: Docker, Docker Compose, Makefile

## Struktur Repositori

```text
.
├── backend/            # Source code untuk Golang backend
├── frontend/           # Source code untuk React/Vite frontend
├── docker-compose.yml  # Pengaturan multi-container Docker untuk environment dev/prod
├── Makefile            # Kumpulan script bantu untuk automasi (build, run, migrate)
├── .env.example        # Contoh environment variables
├── .env.dev            # Environment variables untuk development
├── .env.staging        # Environment variables untuk staging
└── .env.prod           # Environment variables untuk production
```

## Prasyarat (Prerequisites)

Pastikan sistem Anda sudah menginstal perangkat lunak berikut:
1. **Docker** & **Docker Compose**
2. **Git**
3. **Make** (Opsional, sangat disarankan untuk menjalankan script Makefile)
4. **Golang-Migrate CLI** (Opsional, jika ingin menjalankan migrasi tanpa Docker)

---

## 🛠️ Panduan Setup & Instalasi (Dari Awal hingga Production)

### 1. Kloning Repositori & Setup Environment

Pertama, clone repositori ini ke komputer Anda, lalu jalankan script `setup` untuk menduplikasi template environment.

```bash
git clone <url-repo>
cd my-fullstack-app

# Menggandakan .env.example menjadi .env.dev, .env.staging, dan .env.prod
make setup
```

Buka masing-masing file `.env.*` dan sesuaikan parameter koneksi database, kredensial, dan API port sesuai kebutuhan environment Anda.

### 2. Mode Pengembangan (Development)

Untuk menjalankan seluruh layanan (Frontend, Backend, Database) secara instan dalam mode development:

```bash
# Menjalankan kontainer docker dengan .env.dev
make run
```

Atau menggunakan perintah murni tanpa `Make`:
```bash
docker-compose --env-file .env.dev up -d
```

Frontend biasanya akan tersedia di `http://localhost:5173` (atau port lain tergantung definisi Docker Compose), dan Backend API di `http://localhost:8080`.

### 3. Migrasi Database

Terdapat script yang telah disediakan via `Makefile` untuk menangani migrasi database PostgreSQL.

```bash
# Menjalankan migrasi database ke versi terbaru (UP)
make migrate-up

# Membatalkan (rollback) migrasi database (DOWN)
make migrate-down
```
*Catatan*: Pastikan `DB_URL` di dalam `Makefile` telah disesuaikan jika Anda menjalankannya secara native di luar Docker. Di environment production, disarankan menjalankan tools migrasi melalui CI/CD pipeline atau container init eksklusif.

### 4. Build dan Deployment (Production)

Langkah untuk menyiapkan aplikasi agar siap dirilis ke server *Production*:

#### A. Build Ulang Image
Bila Anda melakukan perubahan pada *source code*, bangun ulang Docker image agar menggunakan versi terbaru:

```bash
make build
```

#### B. Jalankan di Mode Production
Gunakan `.env.prod` untuk menjankan production environment. Secara default di Docker Compose, Anda bisa melakukan ini:

```bash
docker-compose --env-file .env.prod up -d
```
Pastikan `docker-compose.yml` Anda juga memiliki konfigurasi spesifik *Production* seperti optimasi stage build Vite menjadi Nginx container, menghilangkan hot-reloading di Go, dan mengoptimalkan resource limits database.

#### C. Memantau Log
Untuk memantau *logs* dari setiap container:

```bash
make logs
# atau
docker-compose logs -f
```

---

## 💻 Script Automasi (Scripts Lengkap)

Selain memanfaatkan `Makefile`, Anda juga bisa menggunakan script manual di bawah ini jika terminal Anda tidak mendukung perintah `make`.

### Skrip Windows (.bat) & Bash (.sh) Opsional
(Bila Anda lebih menyukai skrip batch pada Windows)

**`build_run.bat`**
```bat
@echo off
echo Membangun dan Menjalankan environment Development...
docker-compose --env-file .env.dev up --build -d
```

**`migrate.bat`**
```bat
@echo off
set DB_URL="postgres://myuser:mypassword@localhost:5432/myappdb?sslmode=disable"
migrate -path backend/migrations -database %DB_URL% up
```

## Mematikan (Teardown) Aplikasi

Bila sudah selesai beraktivitas dan ingin mematikan semua layanan:

```bash
# Menghentikan kontainer yang berjalan
make down

# (Peringatan!) Menghapus kontainer beserta volume (mereset seluruh data database)
make clean
```
