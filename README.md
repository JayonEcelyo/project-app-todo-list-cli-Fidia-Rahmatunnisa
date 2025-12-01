📁 To-Do App (CLI)

To-Do App adalah aplikasi manajemen tugas berbasis Command Line Interface (CLI) yang dikembangkan menggunakan Golang. Aplikasi ini membantu pengguna mencatat, melihat, menyelesaikan, mencari, dan menghapus tugas dengan mudah melalui terminal.

Semua data tugas akan tersimpan secara lokal dalam file JSON, sehingga daftar tugas akan aman meskipun aplikasi ditutup.

✨ Fitur Utama
Fitur	Deskripsi
➕ Add Task	Menambahkan tugas baru beserta status & prioritas
📋 List Task	Menampilkan daftar tugas dalam bentuk tabel
✔ Done Task	Mengubah status tugas menjadi completed
❌ Delete Task	Menghapus tugas berdasarkan nomor
🔎 Search Task	Mencari tugas berdasarkan kata kunci
💾 Auto Save	Penyimpanan otomatis ke file JSON
🧱 Struktur Data

Aplikasi menggunakan slice untuk menyimpan data tugas di memori:

type Task struct {
    ID       int    `json:"id"`
    Title    string `json:"title"`
    Status   string `json:"status"`   // new, progress, pending, completed
    Priority string `json:"priority"` // low, medium, high
}

🔧 Fungsi Penting dalam Aplikasi
1️⃣ Menambahkan Tugas

Validasi judul tidak boleh kosong

Mencegah nama tugas duplikat

Status default → new

Data langsung tersimpan ke JSON

2️⃣ Menampilkan Daftar Tugas

Menggunakan fmt.Printf / tabwriter untuk tampilan tabel rapi:

+----+----------------------+------------+----------+
| No | Task                 | Status     | Priority |
+----+----------------------+------------+----------+
| 1  | Cuci                 | completed  | low      |
+----+----------------------+------------+----------+

3️⃣ Menyelesaikan Tugas (Done)

Mengubah status menjadi completed

Validasi ID harus ada dalam daftar

4️⃣ Menghapus Tugas

Menghapus data dari slice menggunakan operasi slice

5️⃣ Load & Save JSON

Menggunakan:

os.ReadFile
os.WriteFile
json.Unmarshal
json.MarshalIndent

🚀 Cara Menggunakan
Clone Repository
git clone https://github.com/JayonEcelyo/project-app-cli-scripting-os-Fidia-Rahmatunnisa.git
cd project-app-todo-list-cli-Fidia-Rahmatunnisa

Menjalankan Command
# Tambah task
go run main.go add --task "belajar golang" --status new --priority high

# List task
go run main.go list

# Set selesai
go run main.go done 1

# Hapus task
go run main.go delete 1

📹 Video Demo

🎥 Demo penggunaan aplikasi:
https://drive.google.com/file/d/1nmG_rDAF4Y4GrBpezdPhWl2JP5kOujPu/view?usp=drive_link

🗂 Struktur Folder
mini-project-2
│
├── cmd
│   ├── add.go
│   ├── delete.go
│   ├── done.go
│   ├── list.go
│   ├── root.go
│
├── data
│   └── tasks.json
│
├── model
│   └── task.go
│
├── service
│   └── task.go
│
├── utils
│   └── storage.go
│
├── go.mod
├── go.sum
└── main.go
