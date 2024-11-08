# Api Dating

## Struktur Projek
```
├───database
│   └───migrations
├───delivery
│   └───http
├───infrastructure
├───middleware
├───models
├───repository
├───testing
├───tmp
├───usecase
└───utils
```

### Project Description
- ```/database```
  - ``` /Migration ```
    
    Database merupakan folder yang berisikan sub-folder yang termasuka di dalam file migration untuk membuat table pada aplikasi api dan mengelola skema database.

- ``` /delivery ```
  - ``` /http ```
  
    Pada lapisan ini Menyediakan rute dan handler untuk representasi api dan komunikasi dengan klien.

- ``` /infrastructure ```
    
    Berisi konfigurasi terkait dengan database dan seeder. database yang dipakai untuk test adalah postgresql versi 11

- ``` /middleware ```

    Menyimpan middleware yang berfungsi sebagai lapisan perantara, memproses permintaan sebelum diteruskan ke handler utama. Contoh middleware: otentikasi, logging, dan validasi.

- ``` /models ```

    Menyimpan definisi model atau struktur data utama yang digunakan dalam api.
- ``` /repository ```

    Lapisan ini berinteraksi dengan tabel pada database dan merupakan dari logika bisnis.

- ``` /testing```

    Folder yang berisi pengujian, pada pengujian ini menggunakan post dan merupakan import file postman untuk testing.

- ```/tmp```

    Folder yang digunakan untuk menyimpan runtime aplikasi menggunakan air golang.

- ``` /usecase ```

    Berisi logika bisnis utama api dan mengimplementasikan berbagai skenario "use case".

- ``` /utils ```

    Berisi fungsi utilitas yang dimana dapat digunakan helper, konversi data, atau fungsi reusable lainnya.

## Menjalankan Migrasi Database

Pada kasus menjalankan proses migrasi ke database menggukan postgresql.
>migrate -database postgres://postgres:Liandi99@localhost/db_dating?sslmode=disable -path database/migrations up

## Menjalankan Project

Pada project ini berjalan di `localhost:4000` dan menggunakan package [air/verse-air](https://github.com/air-verse/air) untuk runtime secara live