# 📝Tugas Besar Alpro2 Kelompok B, Topik SkripsIn
Sistem Informasi Inventaris Dokumen Skripsi adalah aplikasi untuk mengelola arsip tugas
akhir mahasiswa secara digital. Data utama yang digunakan adalah data judul skripsi, data
penulis, dan data tahun lulus. Pengguna aplikasi adalah staf administrasi program studi atau
petugas perpustakaan.
### Anggota Kelompok
1. Nabil Nailur Ridho
2. Joshua Nathanael
3. Firdaus Ramadhana
<br><br>

## 🔨SPESIFIKASI
1. Pengguna dapat menambahkan, mengubah, dan menghapus data dokumen skripsi.
2. Sistem dapat mencatat informasi detail mengenai pembimbing, topik penelitian, dan status kelulusan mahasiswa.
3. Pengguna dapat mencari data skripsi berdasarkan nama mahasiswa atau judul penelitian menggunakan Sequential dan Binary Search.
4. Pengguna dapat mengurutkan data skripsi berdasarkan nama penulis atau tahun lulus menggunakan Selection dan Insertion Sort.
5. Sistem dapat menampilkan statistik jumlah judul skripsi untuk setiap tahun dan total dokumen yang tersimpan.
<br><br>

## 💻MENU UTAMA
1. Tampil Data
2. Tambah Data 
3. Ubah Data  
4. Hapus Data
5. Cari Judul
6. Cari Penulis
7. Urut Nama
8. Urut Tahun
9. Statistik
0. Keluar

<br><br>
## 📖 PENJELASAN MENU UTAMA
1. **Tampil Data**: Menampilkan seluruh daftar dokumen skripsi beserta informasi detail (pembimbing, topik, status) ke layar terminal.
2. **Tambah Data**: Menambahkan data skripsi baru (Judul, Penulis, Tahun, Pembimbing, Topik, Status) ke dalam sistem.
3. **Ubah Data**: Memperbarui informasi skripsi yang sudah ada di dalam memori berdasarkan pencarian ID dokumen.
4. **Hapus Data**: Menghapus data dokumen skripsi dari sistem secara permanen (selama program berjalan) berdasarkan ID dokumen.
5. **Cari Judul (Sequential Search)**: Mencari data spesifik berdasarkan Judul skripsi secara linear (satu per satu dari awal hingga akhir).
6. **Cari Penulis (Binary Search)**: Pencarian data super cepat berdasarkan Nama Penulis (program otomatis mengurutkan abjad A-Z di latar belakang sebelum mencari).
7. **Urut Nama (Selection Sort)**: Mengurutkan daftar dokumen skripsi berdasarkan Nama Penulis secara alfabetis (A-Z).
8. **Urut Tahun (Insertion Sort)**: Alternatif algoritma pengurutan untuk menyusun data dokumen berdasarkan Tahun Lulus (dari yang terlama hingga terbaru).
9. **Statistik**: Menampilkan laporan ringkas berisi total dokumen yang tersimpan dan distribusi jumlah lulusan per tahun.
0. **Keluar**: Menutup aplikasi dengan aman (data *in-memory* akan otomatis dibersihkan).

<br><br>
## 🚀 CARA KERJA (HOW TO RUN THE PROGRAM)
1. Buka folder proyek Anda di aplikasi **Visual Studio Code (VS Code)**.
2. Pastikan file **`main.go`** terbuka di editor.
3. Klik tombol **"Run"** (ikon *play* segitiga) yang ada di pojok kanan atas jendela editor VS Code.
4. Atau jika tidak menggunakan tombol play, buka terminal bawaan dengan klik menu **Terminal > New Terminal**, kemudian ketik perintah: `go run main.go` lalu tekan Enter.
5. Terminal akan otomatis aktif di bagian bawah layar dan program **SkripsIn** siap digunakan!
