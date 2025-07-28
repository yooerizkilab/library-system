# Library System API

REST API untuk sistem manajemen perpustakaan yang dibangun dengan Go, Fiber framework, dan MySQL database.

## 🚀 Features

- **User Management** - CRUD operations untuk pengguna dengan role-based access
- **Book Management** - Manajemen koleksi buku dengan pencarian dan kategorisasi
- **Borrowing System** - Sistem peminjaman buku dengan tracking status
- **Search & Filter** - Pencarian buku dan pengguna
- **Overdue Tracking** - Pelacakan buku yang terlambat dikembalikan
- **RESTful API** - Standard REST endpoints dengan JSON responses
- **Database Relations** - Relational database dengan GORM ORM

## 🛠️ Tech Stack

- **Backend**: Go 1.19+
- **Framework**: Fiber v2
- **Database**: MySQL
- **ORM**: GORM
- **Environment**: godotenv
- **Validation**: go-playground/validator

## 📁 Project Structure

```
library-system-api/
├── cmd/
│   └── server/
│       └── main.go              # Entry point aplikasi
├── internal/
│   ├── config/
│   │   └── config.go            # Konfigurasi environment
│   ├── database/
│   │   └── connection.go        # Database connection
│   ├── models/
│   │   ├── user.go              # User model
│   │   ├── book.go              # Book model
│   │   └── borrow.go            # Borrow model
│   ├── repositories/
│   │   ├── user_repository.go   # User data access
│   │   ├── book_repository.go   # Book data access
│   │   └── borrow_repository.go # Borrow data access
│   ├── services/
│   │   ├── user_service.go      # User business logic
│   │   ├── book_service.go      # Book business logic
│   │   └── borrow_service.go    # Borrow business logic
│   ├── handlers/
│   │   ├── user_handler.go      # User HTTP handlers
│   │   ├── book_handler.go      # Book HTTP handlers
│   │   └── borrow_handler.go    # Borrow HTTP handlers
│   └── routes/
│       └── routes.go            # Route definitions
├── pkg/
│   ├── response/
│   │   └── response.go          # Standardized responses
│   └── utils/
│       └── validator.go         # Utility functions
├── .env                         # Environment variables
├── .env.example                 # Environment template
├── .gitignore                   # Git ignore rules
├── go.mod                       # Go module file
├── go.sum                       # Dependencies checksum
└── README.md                    # Project documentation
```

## 🔧 Installation & Setup

### Prerequisites

- Go 1.19 or higher
- MySQL 5.7 or higher
- Git (optional)

### 1. Clone Repository

```bash
git clone <repository-url>
cd library-system-api
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Setup Database

```sql
-- Login to MySQL
mysql -u root -p

-- Create database
CREATE DATABASE library_system;
```

### 4. Environment Configuration

```bash
# Copy environment template
cp .env.example .env

# Edit environment variables
nano .env
```

Update `.env` file:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_mysql_password
DB_NAME=library_system
APP_PORT=3000
APP_ENV=development
```

### 5. Run Application

```bash
# Development mode
go run cmd/server/main.go

# Or build and run
go build -o bin/library-system-api cmd/server/main.go
./bin/library-system-api
```

Server akan berjalan di `http://localhost:3000`

## 📚 API Documentation

### Base URL

```
http://localhost:3000/api/v1
```

### Authentication

Saat ini API belum menggunakan authentication. Semua endpoints dapat diakses secara bebas.

### Users Endpoints

| Method | Endpoint                | Description     |
| ------ | ----------------------- | --------------- |
| POST   | `/users`                | Create new user |
| GET    | `/users`                | Get all users   |
| GET    | `/users/:id`            | Get user by ID  |
| PUT    | `/users/:id`            | Update user     |
| DELETE | `/users/:id`            | Delete user     |
| GET    | `/users/search?q=query` | Search users    |

### Books Endpoints

| Method | Endpoint                    | Description           |
| ------ | --------------------------- | --------------------- |
| POST   | `/books`                    | Create new book       |
| GET    | `/books`                    | Get all books         |
| GET    | `/books/:id`                | Get book by ID        |
| PUT    | `/books/:id`                | Update book           |
| DELETE | `/books/:id`                | Delete book           |
| GET    | `/books/search?q=query`     | Search books          |
| GET    | `/books/available`          | Get available books   |
| GET    | `/books/category/:category` | Get books by category |

### Borrows Endpoints

| Method | Endpoint                        | Description             |
| ------ | ------------------------------- | ----------------------- |
| POST   | `/borrows`                      | Borrow a book           |
| GET    | `/borrows`                      | Get all borrows         |
| GET    | `/borrows/:id`                  | Get borrow by ID        |
| PUT    | `/borrows/:id`                  | Update borrow           |
| PUT    | `/borrows/:id/return`           | Return a book           |
| GET    | `/borrows/active`               | Get active borrows      |
| GET    | `/borrows/overdue`              | Get overdue borrows     |
| GET    | `/borrows/user/:userId`         | Get user borrows        |
| GET    | `/borrows/book/:bookId`         | Get book borrows        |
| GET    | `/borrows/user/:userId/history` | Get user borrow history |

## 📝 API Examples

### Create User

```bash
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "081234567890",
    "address": "Jl. Example No. 123",
    "role": "member"
  }'
```

### Create Book

```bash
curl -X POST http://localhost:3000/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Clean Code",
    "author": "Robert C. Martin",
    "isbn": "9780132350884",
    "publisher": "Prentice Hall",
    "category": "Programming",
    "language": "English",
    "pages": 464,
    "publish_year": 2008,
    "stock": 3,
    "description": "A handbook of agile software craftsmanship",
    "location": "A1-001"
  }'
```

### Borrow Book

```bash
curl -X POST http://localhost:3000/api/v1/borrows \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "book_id": 1,
    "due_date": "2024-02-15T00:00:00Z",
    "notes": "Borrowed for learning"
  }'
```

## 🎯 Response Format

### Success Response

```json
{
  "status": "success",
  "message": "Operation completed successfully",
  "data": {
    // response data
  }
}
```

### Error Response

```json
{
  "status": "error",
  "message": "Error description",
  "error": "Detailed error information"
}
```

## 🔐 User Roles

- **admin** - Full access to all resources
- **librarian** - Can manage books and borrows
- **member** - Can view books and manage own borrows

## 📊 Borrow Status

- **borrowed** - Book is currently borrowed
- **returned** - Book has been returned
- **overdue** - Book is past due date
- **lost** - Book is reported as lost

## 🧪 Testing

### Manual Testing

Use tools like Postman, Insomnia, or cURL to test the API endpoints.

### Sample Data

Jalankan SQL berikut untuk menambahkan sample data:

```sql
-- Insert sample users
INSERT INTO users (name, email, phone, address, role, is_active, created_at, updated_at) VALUES
('Admin User', 'admin@library.com', '081234567890', 'Jl. Admin No. 1', 'admin', true, NOW(), NOW()),
('John Doe', 'john@example.com', '081234567892', 'Jl. Member No. 3', 'member', true, NOW(), NOW());

-- Insert sample books
INSERT INTO books (title, author, isbn, publisher, category, language, pages, publish_year, stock, available, description, location, is_active, created_at, updated_at) VALUES
('Clean Code', 'Robert C. Martin', '9780132350884', 'Prentice Hall', 'Programming', 'English', 464, 2008, 3, 3, 'A handbook of agile software craftsmanship', 'A1-001', true, NOW(), NOW()),
('Laskar Pelangi', 'Andrea Hirata', '9789793062792', 'Bentang Pustaka', 'Novel', 'Indonesian', 529, 2005, 5, 5, 'Novel tentang pendidikan di Belitung', 'B2-001', true, NOW(), NOW());
```

## 🚨 Troubleshooting

### Database Connection Error

- Pastikan MySQL service berjalan
- Cek kredensial database di file `.env`
- Pastikan database `library_system` sudah dibuat

### Port Already in Use

- Ganti port di `.env` file (APP_PORT)
- Atau stop aplikasi yang menggunakan port tersebut

### Module Not Found

- Jalankan `go mod tidy`
- Pastikan Go version 1.19+

## 🔮 Future Enhancements

- [ ] JWT Authentication & Authorization
- [ ] Role-based access control
- [ ] Input validation middleware
- [ ] Unit & integration tests
- [ ] Rate limiting
- [ ] File upload for book covers
- [ ] Email notifications
- [ ] Admin dashboard
- [ ] Logging system
- [ ] API documentation with Swagger
- [ ] Docker containerization
- [ ] CI/CD pipeline

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Contact

- **Developer**: Your Name
- **Email**: your.email@example.com
- **Project Link**: [https://github.com/yourusername/library-system-api](https://github.com/yourusername/library-system-api)

## 🙏 Acknowledgments

- [Fiber](https://gofiber.io/) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [Go Playground Validator](https://github.com/go-playground/validator) - Validation
- [GoDotEnv](https://github.com/joho/godotenv) - Environment variables
