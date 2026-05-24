#!/bin/env bash

set -e

echo "===================================================="
echo "  Starting E-Money API Deployment Pipeline..."
echo "===================================================="

if [ ! -f .env ]; then
    echo " Error: File .env tidak ditemukan di root directory!"
    exit 1
fi

if [ "$1" == "--reset" ] || [ "$1" == "-r" ]; then
    echo " [Reset Mode] Mematikan service lama dan menghapus database volume..."
    docker compose down -v
    echo " Pembersihan selesai."
fi

echo " Memulai proses kompilasi (build) dan menjalankan container..."
docker compose up --build --remove-orphans -d

echo "===================================================="
echo "         Semua service berhasil dijalankan       "
echo "===================================================="
echo " Status Container:"
docker compose ps

echo -e "\n Menampilkan logs real-time (Tekan Ctrl+C untuk keluar dari logs, aplikasi tetap jalan):"
echo "===================================================="
docker compose logs -f