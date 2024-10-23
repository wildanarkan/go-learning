# 📚 Go CRUD API

> My first project to learn the go language

![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)
![Gin](https://img.shields.io/badge/Gin-v1.9.0-00ADD8?style=flat&logo=go)
![MySQL](https://img.shields.io/badge/MySQL-8.0+-4479A1?style=flat&logo=mysql&logoColor=white)

## 🛠️ Technologi

- **[Go](https://golang.org/)** - Bahasa pemrograman yang cepat dan efisien
- **[Gin](https://gin-gonic.com/)** - Framework web yang ringan dan cepat
- **[GORM](https://gorm.io/)** - ORM yang powerful untuk Go
- **[MySQL](https://www.mysql.com/)** - Database open source yang populer

## 🚀 Quick Start

### Prerequisites

- Go (versi 1.20 atau lebih tinggi)
- MySQL (versi 8.0 atau lebih tinggi)
- Git

### Instalasi

1. **Clone repository**
```bash
git clone https://github.com/wildanarkan/go-learning.git
cd go-learning
```

2. **Install dependencies**
```bash
go mod tidy
```

3. **Setup database**
```sql
CREATE DATABASE go_crud;
```

4. **Jalankan aplikasi**
```bash
go run main.go
```

## 📝 API Endpoints

| Method | Endpoint | Deskripsi |
|--------|----------|------------|
| `GET` | `/api/v1/books` | Mengambil semua buku 📚 |
| `GET` | `/api/v1/books/:id` | Mengambil buku berdasarkan ID 🔍 |
| `POST` | `/api/v1/books` | Menambah buku baru ➕ |
| `PUT` | `/api/v1/books/:id` | Mengupdate buku 📝 |
| `DELETE` | `/api/v1/books/:id` | Menghapus buku ❌ |

## 🙏 Acknowledgments

- [Gin Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [Go Documentation](https://golang.org/doc/)

---

⭐ Project Link: [https://github.com/username/go-crud-api](https://github.com/wildanarkan/go-learning.git)
