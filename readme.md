# Todo App

A full-stack Todo application built with Go (Gin + GORM) backend and Vue.js + TypeScript frontend.

## Features

- User registration and authentication with JWT
- Create, read, update, and delete todo items
- RESTful API with Go/Gin
- MySQL database with GORM ORM
- Vue.js frontend with TypeScript
- Secure user authentication and authorization

## Getting Started

These instructions will help you set up and run the project on your local machine.

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16+)
- [MySQL](https://dev.mysql.com/downloads/) (version 5.7+)
- [Node.js](https://nodejs.org/) (version 14+)
- [npm](https://www.npmjs.com/) or [yarn](https://yarnpkg.com/)

### Clone the Repository

```bash
# Clone the repository
git clone https://github.com/Zhaobo-Wang/go-projects.git

# Navigate to the project directory
cd go-projects/os-go
```

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install Go dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `.env` file by copying `.env.sample`:
   ```bash
   cp .env.sample .env
   ```

4. Update the `.env` file with your MySQL credentials and JWT secret:
   ```
   DB_USER=root
   DB_PASSWORD=your_password
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=todo_app
   JWT_SECRET=your_secret_key
   ```

5. Create the MySQL database:
   ```bash
   mysql -u root -p
   ```
   
   In the MySQL prompt:
   ```sql
   CREATE DATABASE todo_app;
   EXIT;
   ```

6. Run the backend server:
   ```bash
   go run main.go
   ```

   The server will start on http://localhost:8080

## API Endpoints

### Authentication
- `POST /api/v1/register` - Register a new user
- `POST /api/v1/login` - Login and get JWT token

### User
- `GET /api/v1/user` - Get current user information

### Todo Items
- `GET /api/v1/todos` - Get all todos for the authenticated user
- `POST /api/v1/todos` - Create a new todo
- `GET /api/v1/todos/:id` - Get a single todo
- `PUT /api/v1/todos/:id` - Update a todo
- `DELETE /api/v1/todos/:id` - Delete a todo

## Database Schema

### Users Table
- `id` - Primary key
- `username` - Unique username
- `email` - Unique email address
- `password` - Hashed password
- `created_at` - Timestamp
- `updated_at` - Timestamp
- `deleted_at` - Soft delete timestamp

### Todos Table
- `id` - Primary key
- `title` - Todo title
- `description` - Todo description
- `completed` - Boolean status
- `user_id` - Foreign key to users table
- `created_at` - Timestamp
- `updated_at` - Timestamp
- `deleted_at` - Soft delete timestamp

## Built With

- [Go](https://golang.org/) - Backend language
- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [JWT](https://github.com/golang-jwt/jwt) - JSON Web Token authentication
- [Vue.js](https://vuejs.org/) - Frontend framework
- [TypeScript](https://www.typescriptlang.org/) - Type-safe JavaScript
- [MySQL](https://www.mysql.com/) - Database

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [Gin Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [JWT Documentation](https://github.com/golang-jwt/jwt)
- [Vue.js Documentation](https://vuejs.org/guide/introduction.html)