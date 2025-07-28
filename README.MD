# Library System API

REST API untuk sistem manajemen perpustakaan yang dibangun dengan Go, Fiber framework, dan MySQL database.

## 🚀 Features

- **User Management** - CRUD operations untuk pengguna dengan role-based access
- **Authentication & Authorization** - JWT-based login system dengan role-based permissions
- **Book Management** - Manajemen koleksi buku dengan pencarian dan kategorisasi
- **Borrowing System** - Sistem peminjaman buku dengan tracking status
- **Search & Filter** - Pencarian buku dan pengguna
- **Overdue Tracking** - Pelacakan buku yang terlambat dikembalikan
- **Rate Limiting** - API rate limiting (100 requests/minute per IP)
- **RESTful API** - Standard REST endpoints dengan JSON responses
- **Database Relations** - Relational database dengan GORM ORM
- **Password Security** - Bcrypt hashing untuk keamanan password

## 🛠️ Tech Stack

- **Backend**: Go 1.19+
- **Framework**: Fiber v2
- **Database**: MySQL
- **ORM**: GORM
- **Authentication**: JWT (golang-jwt/jwt)
- **Password Hashing**: bcrypt
- **Environment**: godotenv
- **Validation**: go-playground/validator
- **Rate Limiting**: Fiber limiter middleware

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

# Install additional dependencies untuk authentication
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get github.com/gofiber/fiber/v2/middleware/limiter
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
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
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

### 6. Database Migration (Optional)

```bash
# Jalankan SQL migration untuk menambahkan sample users dengan password
mysql -u root -p library_system < migration.sql
```

## 🔐 Authentication

### Default Login Credentials

Setelah menjalankan migration, Anda dapat menggunakan akun default:

```
Admin:
Email: admin@library.com
Password: password123

Librarian:
Email: librarian@library.com
Password: password123

Member:
Email: john@example.com
Password: password123
```

### JWT Token Usage

Untuk mengakses protected endpoints, sertakan JWT token di header:

```bash
Authorization: Bearer <your_jwt_token>
```

## 📚 API Documentation

### Base URL

```
http://localhost:3000/api/v1
```

### Rate Limiting

- Maximum: 100 requests per minute per IP
- Response: 429 Too Many Requests when exceeded

### Authentication Endpoints

| Method | Endpoint         | Description       | Auth Required |
| ------ | ---------------- | ----------------- | ------------- |
| POST   | `/auth/login`    | User login        | No            |
| POST   | `/auth/register` | User registration | No            |

### Profile Endpoints

| Method | Endpoint            | Description      | Auth Required |
| ------ | ------------------- | ---------------- | ------------- |
| GET    | `/profile`          | Get user profile | Yes           |
| PUT    | `/profile/password` | Change password  | Yes           |

### Users Endpoints

| Method | Endpoint                | Description    | Auth Required | Roles            |
| ------ | ----------------------- | -------------- | ------------- | ---------------- |
| GET    | `/users`                | Get all users  | Yes           | Admin, Librarian |
| GET    | `/users/:id`            | Get user by ID | Yes           | Admin, Librarian |
| PUT    | `/users/:id`            | Update user    | Yes           | Admin, Librarian |
| DELETE | `/users/:id`            | Delete user    | Yes           | Admin            |
| GET    | `/users/search?q=query` | Search users   | Yes           | Admin, Librarian |

### Books Endpoints

| Method | Endpoint                    | Description           | Auth Required | Roles            |
| ------ | --------------------------- | --------------------- | ------------- | ---------------- |
| GET    | `/books`                    | Get all books         | No            | Public           |
| GET    | `/books/:id`                | Get book by ID        | No            | Public           |
| GET    | `/books/search?q=query`     | Search books          | No            | Public           |
| GET    | `/books/available`          | Get available books   | No            | Public           |
| GET    | `/books/category/:category` | Get books by category | No            | Public           |
| POST   | `/books/manage`             | Create new book       | Yes           | Admin, Librarian |
| PUT    | `/books/manage/:id`         | Update book           | Yes           | Admin, Librarian |
| DELETE | `/books/manage/:id`         | Delete book           | Yes           | Admin            |

### Borrows Endpoints

| Method | Endpoint                | Description           | Auth Required | Roles            |
| ------ | ----------------------- | --------------------- | ------------- | ---------------- |
| POST   | `/borrows`              | Borrow a book         | Yes           | All              |
| GET    | `/my/borrows`           | Get my borrows        | Yes           | All              |
| GET    | `/my/history`           | Get my borrow history | Yes           | All              |
| GET    | `/borrows/all`          | Get all borrows       | Yes           | Admin, Librarian |
| GET    | `/borrows/:id`          | Get borrow by ID      | Yes           | Admin, Librarian |
| PUT    | `/borrows/:id`          | Update borrow         | Yes           | Admin, Librarian |
| PUT    | `/borrows/:id/return`   | Return a book         | Yes           | Admin, Librarian |
| GET    | `/borrows/active`       | Get active borrows    | Yes           | Admin, Librarian |
| GET    | `/borrows/overdue`      | Get overdue borrows   | Yes           | Admin, Librarian |
| GET    | `/borrows/user/:userId` | Get user borrows      | Yes           | Admin, Librarian |
| GET    | `/borrows/book/:bookId` | Get book borrows      | Yes           | Admin, Librarian |

## 📝 API Examples

### Login

```bash
curl -X POST http://localhost:3000/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@library.com",
    "password": "password123"
  }'
```

### Create User (Registration)

```bash
curl -X POST http://localhost:3000/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "New User",
    "email": "newuser@example.com",
    "password": "password123",
    "phone": "081234567890",
    "address": "Jl. Example No. 123",
    "role": "member"
  }'
```

### Get Profile (Protected)

```bash
curl -X GET http://localhost:3000/api/v1/profile \
  -H "Authorization: Bearer <your_jwt_token>"
```

### Create Book (Protected - Admin/Librarian)

```bash
curl -X POST http://localhost:3000/api/v1/books/manage \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_jwt_token>" \
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

### Borrow Book (Protected)

```bash
curl -X POST http://localhost:3000/api/v1/borrows \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_jwt_token>" \
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

### Authentication Error

- Pastikan JWT token valid dan belum expired
- Cek format Authorization header: `Bearer <token>`
- Pastikan user memiliki role yang sesuai untuk endpoint

### Rate Limit Error

- Tunggu 1 menit sebelum mencoba lagi
- Atau gunakan IP address yang berbeda untuk testing

### Module Not Found

- Jalankan `go mod tidy`
- Pastikan Go version 1.19+

## 🔮 Future Enhancements

- [x] JWT Authentication & Authorization
- [x] Role-based access control
- [x] Rate limiting
- [ ] Token refresh mechanism
- [ ] Input validation middleware
- [ ] Unit & integration tests
- [ ] File upload for book covers
- [ ] Email notifications
- [ ] Admin dashboard
- [ ] Logging system
- [ ] API documentation with Swagger
- [ ] Docker containerization
- [ ] CI/CD pipeline
- [ ] Password reset functionality
- [ ] Account verification via email

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
