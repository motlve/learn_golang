Panduan Dasar Git dan Fungsinya

1. Inisialisasi Repository Git

Perintah:

git init

Fungsi:

Membuat repository Git baru di dalam folder proyek.

Digunakan saat pertama kali ingin menggunakan Git pada suatu proyek.

2. Menambahkan Perubahan ke Staging Area

Perintah:

git add <nama_file>

atau untuk semua file:

git add .

Fungsi:

Memasukkan file ke dalam staging area sebelum dilakukan commit.

Jika menggunakan git add ., maka semua perubahan akan ditambahkan.

3. Menyimpan Perubahan (Commit)

Perintah:

git commit -m "Pesan commit"

Fungsi:

Menyimpan snapshot dari perubahan yang telah ditambahkan ke staging area.

Commit harus memiliki pesan deskriptif tentang perubahan yang dilakukan.

4. Melihat Status Repository

Perintah:

git status

Fungsi:

Menampilkan status file dalam repository.

Menunjukkan file yang telah diubah, ditambahkan, atau belum dipantau oleh Git.

5. Melihat Riwayat Commit

Perintah:

git log

atau versi ringkas:

git log --oneline

Fungsi:

Menampilkan daftar commit yang telah dilakukan.

Versi --oneline menampilkan log dalam format ringkas.

6. Mengecek Perbedaan (Diff) File yang Berubah

Perintah:

git diff

Fungsi:

Menampilkan perubahan yang terjadi pada file sebelum commit.

7. Mengembalikan Perubahan ke Keadaan Sebelumnya

Perintah:

git checkout -- <nama_file>

Fungsi:

Mengembalikan file ke kondisi terakhir di commit atau staging area.

Berguna jika ingin membatalkan perubahan sebelum melakukan commit.

8. Menghapus File dari Repository

Perintah:

git rm <nama_file>

Fungsi:

Menghapus file dari repository dan staging area.

9. Membuat dan Berpindah Branch

Perintah:

git branch <nama_branch>

git checkout <nama_branch>

atau versi singkat:

git switch <nama_branch>

Fungsi:

git branch <nama_branch> → Membuat branch baru.

git checkout <nama_branch> → Berpindah ke branch lain.

git switch <nama_branch> → Alternatif baru untuk berpindah branch.

10. Menggabungkan Branch (Merge)

Perintah:

git merge <nama_branch>

Fungsi:

Menggabungkan branch lain ke dalam branch yang sedang aktif.

11. Menyinkronkan dengan Remote Repository (GitHub/GitLab)

Perintah:

git push origin <nama_branch>

git pull origin <nama_branch>

Fungsi:

git push → Mengunggah commit ke repository online.

git pull → Mengambil perubahan terbaru dari repository online.

12. Mengkloning Repository dari GitHub/GitLab

Perintah:

git clone <url_repository>

Fungsi:

Mendownload repository dari GitHub/GitLab ke komputer lokal.

13. Membatalkan Commit yang Sudah Dilakukan

Perintah:

git reset --soft HEAD~1

atau untuk membatalkan secara permanen:

git reset --hard HEAD~1

Fungsi:

--soft → Membatalkan commit tapi tetap menyimpan perubahan di staging area.

--hard → Menghapus commit beserta perubahan yang dilakukan.