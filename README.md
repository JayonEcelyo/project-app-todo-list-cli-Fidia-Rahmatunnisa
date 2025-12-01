📁 To-Do App (CLI)

To-Do App adalah aplikasi manajemen tugas berbasis Command Line Interface (CLI) yang dikembangkan menggunakan Golang. Aplikasi ini membantu pengguna mencatat, melihat, menyelesaikan, mencari, dan menghapus tugas dengan mudah melalui terminal.

Semua data tugas akan tersimpan secara lokal dalam file JSON, sehingga pengguna dapat membuka kembali daftar tugas meskipun aplikasi ditutup.

✨ Fitur Utama 

Fitur	Deskripsi 

➕ Add Task	Menambahkan tugas baru beserta status & prioritas 

📋 List Task	Menampilkan daftar tugas dalam bentuk tabel 

✔ Done Task	Mengubah status tugas menjadi completed 

❌ Delete Task	Menghapus tugas berdasarkan nomor 

💾 Auto Save	Penyimpanan otomatis ke file JSON 

🧱 Struktur Data 

Aplikasi menggunakan slice untuk menyimpan data tugas di memori. 

type Task struct { 
    ID       int    `json:"id"` 
    Title    string `json:"title"` 
    Status   string `json:"status"`   // new, progress, pending, completed 
    Priority string `json:"priority"` // low, medium, high 
}

🔧 Fungsi Penting pada Aplikasi 
1️⃣ Fungsi Menambahkan Tugas 

Memvalidasi input agar judul tidak kosong 

Menghindari duplikasi nama tugas 

Default status: new 

Menulis ulang JSON ke file 

2️⃣ Fungsi Menampilkan Daftar 

Menggunakan fmt.Printf / tabwriter untuk tampilan tabel rapi: 

+----+----------------------+------------+----------+ 

| No | Task                 | Status     | Priority | 

+----+----------------------+------------+----------+ 

| 1  | Cuci                 | completed  | low      | 

+----+----------------------+------------+----------+ 


3️⃣ Fungsi Update Status (Done) 

Mengubah status menjadi completed 

Validasi ID harus ada 

4️⃣ Fungsi Delete Task 

Menghapus item dari slice menggunakan operasi slice 

5️⃣ Load & Save JSON 

Menggunakan: 

os.WriteFile, os.ReadFile 
json.Unmarshal, json.MarshalIndent 

🚀 Cara Menggunakan 
Clone repo 
git clone https://github.com/JayonEcelyo/project-app-cli-scripting-os-Fidia-Rahmatunnisa.git 
cd project-app-todo-list-cli-Fidia-Rahmatunnisa 
[add task]: go run main.go add --task "belajar golang" --status new --priority high 
[list task]: go run main.go list 
[done task]: go run main.go done ID (example : go run main.go done 1) 
[delete task]: go run main.go delete ID (example : go run main.go delete 1) 

📹 Video Demo 
🎥 Demo penggunaan aplikasi: https://drive.google.com/file/d/1nmG_rDAF4Y4GrBpezdPhWl2JP5kOujPu/view?usp=drive_link  

🗂 Struktur Folder 
mini-project-2 
│  
├── cmd   
  ├── add.go  
  ├── delete.go   
  ├── done.go   
  ├── list.go   
  ├── root.go   
├── data   
  ├── tasks.json   
├── model   
  ├── task.go   
├── service   
  ├── task.go   
├── utils   
  ├── storage.go   
└── go.mod   
└── go.sum   
└── main.go  
